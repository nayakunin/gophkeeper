package register

import (
	"fmt"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

// Service is the register service.
func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			encryptionKey, err := s.encryption.GenerateKey()

			if err != nil {
				return fmt.Errorf("could not generate encryption key: %w", err)
			}

			response, err := s.api.RegisterUser(cmd.Context(), &api.RegisterUserRequest{
				Username: username,
				Password: password,
			})
			if err != nil {
				return fmt.Errorf("could not register user: %w", err)
			}

			if err := s.storage.SaveCredentials(response.GetToken(), encryptionKey); err != nil {
				return fmt.Errorf("could not save credentials: %w", err)
			}

			fmt.Println("Successfully registered and logged in")
			return nil
		},
	}

	cmd.Flags().StringP("username", "u", "", "Username for the new user")
	cmd.Flags().StringP("password", "p", "", "Password for the new user")

	return cmd
}
