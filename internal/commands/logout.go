package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for logging out
		fmt.Println("Logging out...")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
