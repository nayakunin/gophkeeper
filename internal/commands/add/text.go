package add

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) textCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Add a new text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)

			text, err := cmd.Flags().GetString("text")
			if err != nil {
				return fmt.Errorf("could not get text: %w", err)
			}
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("could not get description: %w", err)
			}

			encryptedText, err := encryption.Encrypt(text, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not encrypt text: %w", err)
			}

			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			_, err = client.AddTextData(ctx, &api.AddTextDataRequest{
				Description:   description,
				EncryptedText: encryptedText,
			})
			if err != nil {
				return fmt.Errorf("could not add text data: %w", err)
			}

			fmt.Println("Successfully added text data")
			return nil
		},
	}

	cmd.Flags().StringP("text", "t", "", "Text content")
	_ = cmd.MarkFlagRequired("text")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
