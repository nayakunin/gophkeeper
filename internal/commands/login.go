package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/nayakunin/gophkeeper/constants"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()

		client := api.NewAuthServiceClient(conn)

		response, err := client.AuthenticateUser(context.Background(), &api.AuthenticateUserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalf("Could not authenticate user: %v", err)
		}

		fmt.Printf("Authentication result: %s\n", response.String())
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	// Add more flags if needed

	rootCmd.AddCommand(loginCmd)
}
