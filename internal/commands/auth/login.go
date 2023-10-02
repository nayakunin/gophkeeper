package auth

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func (s *Service) LoginCmd() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Log in as a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewAuthServiceClient(conn)

			response, err := client.AuthenticateUser(cmd.Context(), &api.AuthenticateUserRequest{
				Username: username,
				Password: []byte(password),
			})
			if err != nil {
				return fmt.Errorf("could not authenticate user: %w", err)
			}

			if err := s.storage.SaveCredentials(response.GetToken(), response.GetEncryptionKey()); err != nil {
				return fmt.Errorf("could not save credentials: %w", err)
			}

			fmt.Println("Successfully logged in")
			return nil
		},
	}

	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	_ = loginCmd.MarkFlagRequired("username")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	_ = loginCmd.MarkFlagRequired("password")

	return loginCmd
}
