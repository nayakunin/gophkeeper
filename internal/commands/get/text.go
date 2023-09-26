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

func (s *Service) textCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Get a text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}
			if token == "" {
				return fmt.Errorf("please login first")
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)

			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			response, err := client.GetTextData(ctx, &api.Empty{})
			if err != nil {
				return fmt.Errorf("could not get text data: %w", err)
			}

			type Result struct {
				Description string `json:"description"`
				Text        string `json:"text"`
			}

			results := make([]Result, len(response.GetTextData()))
			for i, textData := range response.GetTextData() {
				decryptedText, err := s.encryption.Decrypt(textData.GetEncryptedText(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt text: %w", err)
				}

				results[i] = Result{
					Description: textData.GetDescription(),
					Text:        decryptedText,
				}
			}

			return utils.PrintJSON(results)
		},
	}

	return cmd
}
