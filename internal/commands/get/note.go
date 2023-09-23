package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NoteCmd = &cobra.Command{
	Use:   "note",
	Short: "Get a note",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for getting a password
		fmt.Println("Getting note...")
	},
}
