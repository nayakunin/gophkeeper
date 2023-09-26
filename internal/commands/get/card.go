package get

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) cardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Get a card",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}
			if token == "" {
				return fmt.Errorf("please login first")
			}

			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("could not get card name: %w", err)
			}
			if name == "" {
				return fmt.Errorf("please provide a card name")
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			response, err := client.GetBankCardDetails(ctx, &api.GetBankCardDetailsRequest{
				CardName: name,
			})
			if err != nil {
				return fmt.Errorf("could not get card: %w", err)
			}

			type Result struct {
				Name        string `json:"name"`
				Number      string `json:"number"`
				Expiration  string `json:"expiration"`
				Cvc         string `json:"cvv"`
				Description string `json:"description"`
			}
			results := make([]Result, len(response.GetBankCardDetails()))
			for i, card := range response.GetBankCardDetails() {
				number, err := s.encryption.Decrypt(card.GetEncryptedCardNumber(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt card number: %w", err)
				}
				expiration, err := s.encryption.Decrypt(card.GetEncryptedExpiryDate(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt card expiration date: %w", err)
				}
				cvc, err := s.encryption.Decrypt(card.GetEncryptedCvc(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt card CVC: %w", err)
				}
				results[i] = Result{
					Name:        card.GetCardName(),
					Number:      number,
					Expiration:  expiration,
					Cvc:         cvc,
					Description: card.GetDescription(),
				}
			}

			return utils.PrintJSON(results)
		},
	}

	cmd.Flags().StringP("name", "n", "", "Card name")

	return cmd
}
