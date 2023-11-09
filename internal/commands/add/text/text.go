package text

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	"github.com/spf13/cobra"
)

// GetCmd returns the text command.
func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Add a new text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}
			s.api.SetToken(token)

			tmpResult, err := input.ParseTextRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PrepareTextRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			err = s.api.AddTextData(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("could not add text data: %w", err)
			}

			fmt.Println("Successfully added text data")
			return nil
		},
	}

	cmd.Flags().StringP("text", "t", "", "Text content")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
