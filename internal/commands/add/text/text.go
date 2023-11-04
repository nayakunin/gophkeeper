package text

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "text",
		Short: "Add a new text data",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := input.ParseTextRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PrepareTextRequest(tmpResult, encryptionKey)
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
