package commands

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("could not connect: %w", err)
		}
		defer conn.Close()

		client := api.NewRegistrationServiceClient(conn)

		response, err := client.RegisterUser(context.Background(), &api.RegisterUserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			return fmt.Errorf("could not register user: %w", err)
		}

		fmt.Printf("Registration result: %s\n", response.GetMessage())
		return nil
	},
}

func init() {
	registerCmd.Flags().StringP("username", "u", "", "Username for the new user")
	registerCmd.Flags().StringP("password", "p", "", "Password for the new user")
	rootCmd.AddCommand(registerCmd)
}
