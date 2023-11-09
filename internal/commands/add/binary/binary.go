package binary

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	"github.com/spf13/cobra"
)

// GetCmd returns the binary command.
func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Add a new binary record",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}
			s.api.SetToken(token)

			tmpResult, err := input.ParseBinaryRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PrepareBinaryRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			err = s.api.AddBinaryData(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("could not add binary data: %w", err)
			}

			fmt.Println("Binary data added successfully")
			return nil
		},
	}

	cmd.Flags().StringP("filepath", "f", "", "File path")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
