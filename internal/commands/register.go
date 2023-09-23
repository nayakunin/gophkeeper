package commands

import (
	"context"
	"fmt"
	"log"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()

		client := api.NewUserServiceClient(conn)

		response, err := client.RegisterUser(context.Background(), &api.RegisterUserRequest{
			Username: username,
			Email:    email,
			Password: password,
		})
		if err != nil {
			log.Fatalf("Could not register user: %v", err)
		}

		fmt.Printf("Registration result: %s\n", response.GetMessage())
	},
}

func init() {
	registerCmd.Flags().StringP("username", "u", "", "Username for the new user")
	registerCmd.Flags().StringP("email", "e", "", "Email address for the new user")
	registerCmd.Flags().StringP("password", "p", "", "Password for the new user")
	rootCmd.AddCommand(registerCmd)
}
