package add

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type parseTextResult struct {
	Text        string
	Description string
}

func (s *Service) parseTextRequest(cmd *cobra.Command) (*parseTextResult, error) {
	text, err := cmd.Flags().GetString("text")
	if err != nil {
		return nil, fmt.Errorf("could not get text: %w", err)
	}
	if text == "" {
		return nil, fmt.Errorf("please provide a text")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &parseTextResult{
		Text:        text,
		Description: description,
	}, nil
}

func (s *Service) prepareTextRequest(result *parseTextResult, encryptionKey []byte) (*api.AddTextDataRequest, error) {
	encryptedText, err := s.encryption.Encrypt([]byte(result.Text), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt text: %w", err)
	}

	return &api.AddTextDataRequest{
		EncryptedText: encryptedText,
		Description:   result.Description,
	}, nil
}

func (s *Service) textCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Add a new text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := s.parseTextRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.prepareTextRequest(tmpResult, encryptionKey)
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
			ctx := metadata.NewOutgoingContext(cmd.Context(), md)
			_, err = client.AddTextData(ctx, request)
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
