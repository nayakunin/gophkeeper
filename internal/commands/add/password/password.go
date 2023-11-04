package password

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Add a new password",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := input.ParsePasswordRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PreparePasswordRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(cmd.Context(), md)
			_, err = client.AddLoginPasswordPair(ctx, request)
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
