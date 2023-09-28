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

func (s *Service) binaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Get a binary record",
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
			response, err := client.GetBinaryData(ctx, &api.Empty{})
			if err != nil {
				return fmt.Errorf("could not get binary data: %w", err)
			}

			type Result struct {
				Data        []byte `json:"data"`
				Description string `json:"description"`
			}
			results := make([]Result, len(response.GetBinaryData()))
			for i, pair := range response.GetBinaryData() {
				data, err := s.encryption.Decrypt(pair.GetEncryptedData(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt data: %w", err)
				}
				results[i] = Result{
					Data:        data,
					Description: pair.GetDescription(),
				}
			}

			return utils.PrintJSON(results)
		},
	}

	return cmd
}
