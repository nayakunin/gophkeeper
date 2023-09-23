package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var BinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Get a binary record",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for getting a password
		fmt.Println("Getting binary record...")
	},
}
