package get

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get an entry",
}

func init() {
	Cmd.AddCommand(PasswordCmd)
	Cmd.AddCommand(CardCmd)
	Cmd.AddCommand(BinaryCmd)
	Cmd.AddCommand(NoteCmd)
}
