package binary

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/spf13/cobra"
)

// GetCmd returns the binary command.
func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Get a binary record",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}
			s.api.SetToken(token)

			response, err := s.api.GetBinaryData(cmd.Context())
			if err != nil {
				return fmt.Errorf("could not get binary data: %w", err)
			}

			results, err := s.output.MakeResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make binary response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	} 

	return cmd
}
