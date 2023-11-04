package card

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) GetCardCmd() *cobra.Command {
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
