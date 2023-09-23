package commands

import (
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in as a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Your logic to login.
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username for login")
	loginCmd.Flags().StringP("password", "p", "", "Password for login")
	// Add more flags if needed

	rootCmd.AddCommand(loginCmd)
}
