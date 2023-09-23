package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NoteCmd = &cobra.Command{
	Use:   "note",
	Short: "Add a new note",
	Run: func(cmd *cobra.Command, args []string) {
		// Logic for adding a password
		fmt.Println("Adding note data...")
	},
}
