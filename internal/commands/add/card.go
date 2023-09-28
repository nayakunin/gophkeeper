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

			name, err := cmd.Flags().GetString("label")
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
			cvc, err := cmd.Flags().GetString("cvc")
			if err != nil {
				return fmt.Errorf("could not get card CVV: %w", err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {

			}

			encryptedNumber, err := s.encryption.Encrypt([]byte(number), encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card number: %w", err)
			}
			encryptedExpiration, err := s.encryption.Encrypt([]byte(expiration), encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card expiration date: %w", err)
			}
			encryptedCVC, err := s.encryption.Encrypt([]byte(cvc), encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt card CVC: %w", err)
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
				EncryptedCvc:        encryptedCVC,
				Description:         description,
			})
			if err != nil {
				return fmt.Errorf("could not add card: %w", err)
			}

			fmt.Println("Card added")
			return nil
		},
	}

	cmd.Flags().StringP("label", "l", "", "Card label")
	_ = cmd.MarkFlagRequired("name")
	cmd.Flags().StringP("number", "n", "", "Card number")
	_ = cmd.MarkFlagRequired("number")
	cmd.Flags().StringP("expiration", "e", "", "Card expiration date")
	_ = cmd.MarkFlagRequired("expiration")
	cmd.Flags().StringP("cvc", "c", "", "Card CVC")
	_ = cmd.MarkFlagRequired("cvc")
	cmd.Flags().StringP("description", "d", "", "Card description")

	return cmd
}
