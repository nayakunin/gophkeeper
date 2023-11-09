package card

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
	"github.com/spf13/cobra"
)

// GetCmd returns the card command.
func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Add a new credit card",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := input.ParseCardRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PrepareCardRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			s.api.SetToken(token)
			err = s.api.AddCardData(context.Background(), request)
			if err != nil {
				return fmt.Errorf("could not add card data: %w", err)
			}

			fmt.Println("Card added")
			return nil
		},
	}

	cmd.Flags().StringP("label", "l", "", "Card label")
	cmd.Flags().StringP("number", "n", "", "Card number")
	cmd.Flags().StringP("expiration", "e", "", "Card expiration date")
	cmd.Flags().StringP("cvc", "c", "", "Card CVC")
	cmd.Flags().StringP("description", "d", "", "Card description")

	return cmd
}
