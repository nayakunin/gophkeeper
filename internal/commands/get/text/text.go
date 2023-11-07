package text

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/spf13/cobra"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Get a text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}

			response, err := s.api.GetTextData(cmd.Context())
			if err != nil {
				return fmt.Errorf("could not get text data: %w", err)
			}

			results, err := s.output.MakeResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make text response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	}

	return cmd
}
