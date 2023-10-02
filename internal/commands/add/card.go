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

type parseCardResult struct {
	Name        string
	Number      string
	Expiration  string
	Cvc         string
	Description string
}

func (s *Service) parseCardRequest(cmd *cobra.Command) (*parseCardResult, error) {
	name, err := cmd.Flags().GetString("label")
	if err != nil {
		return nil, fmt.Errorf("could not get card name: %w", err)
	}
	if name == "" {
		return nil, fmt.Errorf("please provide a card name")
	}
	number, err := cmd.Flags().GetString("number")
	if err != nil {
		return nil, fmt.Errorf("could not get card number: %w", err)
	}
	if number == "" {
		return nil, fmt.Errorf("please provide a card number")
	}
	expiration, err := cmd.Flags().GetString("expiration")
	if err != nil {
		return nil, fmt.Errorf("could not get card expiration date: %w", err)
	}
	if expiration == "" {
		return nil, fmt.Errorf("please provide a card expiration date")
	}
	cvc, err := cmd.Flags().GetString("cvc")
	if err != nil {
		return nil, fmt.Errorf("could not get card CVC: %w", err)
	}
	if cvc == "" {
		return nil, fmt.Errorf("please provide a card CVC")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get card description: %w", err)
	}

	return &parseCardResult{
		Name:        name,
		Number:      number,
		Expiration:  expiration,
		Cvc:         cvc,
		Description: description,
	}, nil
}

func (s *Service) prepareCardRequest(data *parseCardResult, encryptionKey []byte) (*api.AddBankCardDetailRequest, error) {
	encryptedNumber, err := s.encryption.Encrypt([]byte(data.Number), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card number: %w", err)
	}
	encryptedExpiration, err := s.encryption.Encrypt([]byte(data.Expiration), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card expiration date: %w", err)
	}
	encryptedCVC, err := s.encryption.Encrypt([]byte(data.Cvc), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt card CVC: %w", err)
	}

	return &api.AddBankCardDetailRequest{
		CardName:            data.Name,
		EncryptedCardNumber: encryptedNumber,
		EncryptedExpiryDate: encryptedExpiration,
		EncryptedCvc:        encryptedCVC,
		Description:         data.Description,
	}, nil
}

func (s *Service) cardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Add a new credit card",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := s.parseCardRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.prepareCardRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			_, err = client.AddBankCardDetail(ctx, request)
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
