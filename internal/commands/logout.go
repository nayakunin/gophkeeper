package commands

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/credentials"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := credentials.Store.Delete("token")
		if err != nil {
			return fmt.Errorf("unable to logout: %w", err)
		}

		fmt.Println("Successfully logged out")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
