package add

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) passwordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Add a new password",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			serviceName, err := cmd.Flags().GetString("service")
			if err != nil {
				return fmt.Errorf("could not get service name: %w", err)
			}
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				return fmt.Errorf("could not get login: %w", err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return fmt.Errorf("could not get password: %w", err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("could not get description: %w", err)
			}

			encryptedPassword, err := s.encryption.Encrypt([]byte(password), encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt password: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			_, err = client.AddLoginPasswordPair(ctx, &api.AddLoginPasswordPairRequest{
				ServiceName:       serviceName,
				Login:             login,
				EncryptedPassword: encryptedPassword,
				Description:       description,
			})
			if err != nil {
				return fmt.Errorf("could not add password: %w", err)
			}

			fmt.Println("Password added")
			return nil
		},
	}

	cmd.Flags().StringP("service", "s", "", "Service name")
	_ = cmd.MarkFlagRequired("service")
	cmd.Flags().StringP("login", "l", "", "Login")
	_ = cmd.MarkFlagRequired("login")
	cmd.Flags().StringP("password", "p", "", "Password")
	_ = cmd.MarkFlagRequired("password")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
