package add

import (
	"context"
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
	Short: "Add a new password",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := credentials.Store.Get("token")
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

		serviceName, err := cmd.Flags().GetString("service")
		if err != nil {
			return fmt.Errorf("could not get service name: %w", err)
		}
		if serviceName == "" {
			return fmt.Errorf("please provide a service name")
		}
		login, err := cmd.Flags().GetString("login")
		if err != nil {
			return fmt.Errorf("could not get login: %w", err)
		}
		if login == "" {
			return fmt.Errorf("please provide a login")
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return fmt.Errorf("could not get password: %w", err)
		}
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			return fmt.Errorf("could not get description: %w", err)
		}

		fmt.Println(token)

		md := credentials.GetRequestMetadata(token)
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		response, err := client.AddLoginPasswordPair(ctx, &api.AddLoginPasswordPairRequest{
			ServiceName: serviceName,
			Login:       login,
			Password:    password,
			Description: description,
		})
		if err != nil {
			return fmt.Errorf("could not add password: %w", err)
		}

		fmt.Printf("Password added: %s\n", response.GetMessage())

		return nil
	},
}

func init() {
	PasswordCmd.Flags().StringP("service", "s", "", "Service name")
	PasswordCmd.Flags().StringP("login", "l", "", "Login")
	PasswordCmd.Flags().StringP("password", "p", "", "Password")
	PasswordCmd.Flags().StringP("description", "d", "", "Description")
}
