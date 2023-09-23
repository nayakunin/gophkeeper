package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CardCmd = &cobra.Command{
	Use:   "card",
	Short: "Get a card",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for getting a password
		fmt.Println("Getting card...")
	},
}
