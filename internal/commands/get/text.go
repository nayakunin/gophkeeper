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

type textResult struct {
	Description string `json:"description"`
	Text        string `json:"text"`
}

func (s *Service) makeTextResponse(response *api.GetTextDataResponse, encryptionKey []byte) ([]textResult, error) {
	results := make([]textResult, len(response.GetTextData()))
	for i, textData := range response.GetTextData() {
		decryptedText, err := s.encryption.Decrypt(textData.GetEncryptedText(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt text: %w", err)
		}

		results[i] = textResult{
			Description: textData.GetDescription(),
			Text:        string(decryptedText),
		}
	}

	return results, nil
}

func (s *Service) textCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Get a text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
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

			results, err := s.makeTextResponse(response, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not make text response: %w", err)
			}

			return utils.PrintJSON(results)
		},
	}

	return cmd
}
