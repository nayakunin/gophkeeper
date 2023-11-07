package password

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/spf13/cobra"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Get a password",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}

			request, err := input.ParsePasswordRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			response, err := s.api.GetLoginPasswordPairs(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("could not get password: %w", err)
			}

			results, err := s.output.MakeResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make password response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	}

	cmd.Flags().StringP("service", "s", "", "Service name")

	return cmd
}
