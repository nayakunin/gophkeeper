package get

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) parseCardRequest(cmd *cobra.Command) (*api.GetBankCardDetailsRequest, error) {
	label, err := cmd.Flags().GetString("label")
	if err != nil {
		return nil, fmt.Errorf("could not get card label: %w", err)
	}
	if label == "" {
		return nil, fmt.Errorf("please provide a card label")
	}

	return &api.GetBankCardDetailsRequest{
		CardName: label,
	}, nil
}

type cardResult struct {
	Name        string `json:"label"`
	Number      string `json:"number"`
	Expiration  string `json:"expiration"`
	Cvc         string `json:"cvv"`
	Description string `json:"description"`
}

func (s *Service) makeCardResponse(response *api.GetBankCardDetailsResponse, encryptionKey []byte) ([]cardResult, error) {
	results := make([]cardResult, len(response.GetBankCardDetails()))
	for i, card := range response.GetBankCardDetails() {
		number, err := s.encryption.Decrypt(card.GetEncryptedCardNumber(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card number: %w", err)
		}
		expiration, err := s.encryption.Decrypt(card.GetEncryptedExpiryDate(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card expiration date: %w", err)
		}
		cvc, err := s.encryption.Decrypt(card.GetEncryptedCvc(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card CVC: %w", err)
		}
		results[i] = cardResult{
			Name:        card.GetCardName(),
			Number:      string(number),
			Expiration:  string(expiration),
			Cvc:         string(cvc),
			Description: card.GetDescription(),
		}
	}

	return results, nil
}

func (s *Service) cardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Get a card",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}

			request, err := s.parseCardRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse card request: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(cmd.Context(), md)
			response, err := client.GetBankCardDetails(ctx, request)
			if err != nil {
				return fmt.Errorf("could not get card: %w", err)
			}

			results, err := s.makeCardResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make card response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	}

	cmd.Flags().StringP("label", "l", "", "Card label")

	return cmd
}
