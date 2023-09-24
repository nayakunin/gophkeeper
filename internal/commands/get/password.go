package get

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/credentials"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var PasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Get a password",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := credentials.Store.Get("token")
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

		md := credentials.GetRequestMetadata(token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		response, err := client.GetLoginPasswordPairs(ctx, &api.GetLoginPasswordPairsRequest{
			ServiceName: serviceName,
		})
		if err != nil {
			return fmt.Errorf("could not get password: %w", err)
		}

		jsonData, err := json.MarshalIndent(response.GetLoginPasswordPairs(), "", "  ")
		if err != nil {
			return fmt.Errorf("could not marshal response: %w", err)
		}

		fmt.Printf("Passwords for %s: %s\n", serviceName, string(jsonData))
		return nil
	},
}

func init() {
	PasswordCmd.Flags().StringP("service", "s", "", "Service name")
}
