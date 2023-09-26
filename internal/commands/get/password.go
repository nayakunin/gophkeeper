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

func (s *Service) passwordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Get a password",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get token: %w", err)
			}
			if token == "" {
				return fmt.Errorf("please login first")
			}

			serviceName, err := cmd.Flags().GetString("service")
			if err != nil {
				return fmt.Errorf("could not get service name: %w", err)
			}
			if serviceName == "" {
				return fmt.Errorf("please provide a service name")
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)

			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			response, err := client.GetLoginPasswordPairs(ctx, &api.GetLoginPasswordPairsRequest{
				ServiceName: serviceName,
			})
			if err != nil {
				return fmt.Errorf("could not get password: %w", err)
			}

			type Result struct {
				ServiceName string `json:"service_name"`
				Login       string `json:"login"`
				Password    string `json:"password"`
				Description string `json:"description"`
			}
			results := make([]Result, len(response.GetLoginPasswordPairs()))
			for i, pair := range response.GetLoginPasswordPairs() {
				password, err := s.encryption.Decrypt(pair.GetEncryptedPassword(), encryptionKey)
				if err != nil {
					return fmt.Errorf("could not decrypt password: %w", err)
				}

				results[i] = Result{
					ServiceName: pair.GetServiceName(),
					Login:       pair.GetLogin(),
					Password:    password,
					Description: pair.GetDescription(),
				}
			}

			return utils.PrintJSON(results)
		},
	}

	cmd.Flags().StringP("service", "s", "", "Service name")

	return cmd
}
