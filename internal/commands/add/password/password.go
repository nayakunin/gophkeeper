package password

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	"github.com/spf13/cobra"
)

// GetCmd returns the password command.
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

			s.api.SetToken(token)
			err = s.api.AddPasswordData(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("could not add password: %w", err)
			}

			fmt.Println("Password added")
			return nil
		},
	}

	cmd.Flags().StringP("service", "s", "", "Service name")
	cmd.Flags().StringP("login", "l", "", "Login")
	cmd.Flags().StringP("password", "p", "", "Password")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
