package card

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/get/card/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/spf13/cobra"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Get a card",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}

			request, err := input.ParseCardRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse card request: %w", err)
			}

			response, err := s.api.GetCardDetails(cmd.Context(), request)
			if err != nil {
				return fmt.Errorf("could not get card: %w", err)
			}

			results, err := s.output.MakeResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make card response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	}

	cmd.Flags().StringP("label", "l", "", "Card label")

	return cmd
}
