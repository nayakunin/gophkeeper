package auth

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func (s *Service) RegisterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewRegistrationServiceClient(conn)

			encryptionKey, err := encryption.GenerateKey()
			fmt.Println("encryptionKey: ", hex.EncodeToString(encryptionKey))

			if err != nil {
				return fmt.Errorf("could not generate encryption key: %w", err)
			}

			response, err := client.RegisterUser(context.Background(), &api.RegisterUserRequest{
				Username: username,
				Password: password,
			})
			if err != nil {
				return fmt.Errorf("could not register user: %w", err)
			}

			if err := s.storage.SaveCredentials(response.GetToken(), hex.EncodeToString(encryptionKey)); err != nil {
				return fmt.Errorf("could not save credentials: %w", err)
			}

			fmt.Println("Successfully registered and logged in")
			return nil
		},
	}

	cmd.Flags().StringP("username", "u", "", "Username for the new user")
	_ = cmd.MarkFlagRequired("username")
	cmd.Flags().StringP("password", "p", "", "Password for the new user")
	_ = cmd.MarkFlagRequired("password")

	return cmd
}
