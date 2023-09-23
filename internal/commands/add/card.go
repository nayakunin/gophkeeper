package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CardCmd = &cobra.Command{
	Use:   "card",
	Short: "Add a new credit card",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for adding a password
		fmt.Println("Adding credit card...")
	},
}
