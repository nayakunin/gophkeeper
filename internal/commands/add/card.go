package add

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
		Short: "Add a new credit card",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("could not get card name: %w", err)
			}
			number, err := cmd.Flags().GetString("number")
			if err != nil {
				return fmt.Errorf("could not get card number: %w", err)
			}
			expiration, err := cmd.Flags().GetString("expiration")
			if err != nil {
				return fmt.Errorf("could not get card expiration date: %w", err)
			}
			cvv, err := cmd.Flags().GetString("cvv")
			if err != nil {
				return fmt.Errorf("could not get card CVV: %w", err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {

			}

			encryptedNumber, err := s.encryption.Encrypt(number, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card number: %w", err)
			}
			encryptedExpiration, err := s.encryption.Encrypt(expiration, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card expiration date: %w", err)
			}
			encryptedCvv, err := s.encryption.Encrypt(cvv, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card CVV: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			_, err = client.AddBankCardDetail(ctx, &api.AddBankCardDetailRequest{
				CardName:            name,
				EncryptedCardNumber: encryptedNumber,
				EncryptedExpiryDate: encryptedExpiration,
				EncryptedCvc:        encryptedCvv,
				Description:         description,
			})
			if err != nil {
				return fmt.Errorf("could not add card: %w", err)
			}

			fmt.Println("Card added")
			return nil
		},
	}

	cmd.Flags().String("name", "", "Card name")
	_ = cmd.MarkFlagRequired("name")
	cmd.Flags().String("number", "", "Card number")
	_ = cmd.MarkFlagRequired("number")
	cmd.Flags().String("expiration", "", "Card expiration date")
	_ = cmd.MarkFlagRequired("expiration")
	cmd.Flags().String("cvv", "", "Card CVV")
	_ = cmd.MarkFlagRequired("cvv")
	cmd.Flags().String("description", "", "Card description")

	return cmd
}
