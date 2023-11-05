package login

import (
	"fmt"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

func (s *Service) GetCmd() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Log in as a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			response, err := s.api.AuthenticateUser(cmd.Context(), &api.AuthenticateUserRequest{
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
	loginCmd.Flags().StringP("password", "p", "", "Password for login")

	return loginCmd
}
