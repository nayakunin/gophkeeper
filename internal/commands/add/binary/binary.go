package binary

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (s *Service) GetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Add a new binary record",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := input.ParseBinaryRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.apiPreparer.PrepareBinaryRequest(tmpResult, encryptionKey)
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
			_, err = client.AddBinaryData(ctx, request)
			if err != nil {
				return fmt.Errorf("could not add binary data: %w", err)
			}

			fmt.Println("Binary data added successfully")
			return nil
		},
	}

	cmd.Flags().StringP("filepath", "f", "", "File path")
	_ = cmd.MarkFlagRequired("filepath")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
