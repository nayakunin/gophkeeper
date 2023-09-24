package commands

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/credentials"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("could not connect: %w", err)
		}
		defer conn.Close()

		client := api.NewAuthServiceClient(conn)

		response, err := client.AuthenticateUser(context.Background(), &api.AuthenticateUserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return fmt.Errorf("could not authenticate user: %w", err)
		}

		if err := credentials.Store.Set("token", response.GetToken()); err != nil {
			return fmt.Errorf("could not set token: %w", err)
		}

		fmt.Println("Successfully logged in")
		return nil
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	// Add more flags if needed

	rootCmd.AddCommand(loginCmd)
}
