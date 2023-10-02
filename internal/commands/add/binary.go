package add

import (
	"fmt"
	"os"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type parseBinaryResult struct {
	Filepath    string
	Description string
}

func (s *Service) parseBinaryRequest(cmd *cobra.Command, encryptionKey []byte) (*parseBinaryResult, error) {
	filepath, err := cmd.Flags().GetString("filepath")
	if err != nil {
		return nil, fmt.Errorf("could not get filepath: %w", err)
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &parseBinaryResult{
		Filepath:    filepath,
		Description: description,
	}, nil
}

func (s *Service) prepareBinaryRequest(result *parseBinaryResult, encryptionKey []byte) (*api.AddBinaryDataRequest, error) {
	file, err := os.ReadFile(result.Filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	encryptedFile, err := s.encryption.Encrypt(file, encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt file: %w", err)
	}

	return &api.AddBinaryDataRequest{
		EncryptedData: encryptedFile,
		Description:   result.Description,
	}, nil
}

func (s *Service) binaryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Add a new binary record",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := s.parseBinaryRequest(cmd, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.prepareBinaryRequest(tmpResult, encryptionKey)
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
