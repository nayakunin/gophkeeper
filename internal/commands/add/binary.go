package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var BinaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Add a new binary record",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for adding a password
		fmt.Println("Adding binary data...")
	},
}
