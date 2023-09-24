package add

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
}

func init() {
	Cmd.AddCommand(PasswordCmd)
	Cmd.AddCommand(CardCmd)
	Cmd.AddCommand(BinaryCmd)
	Cmd.AddCommand(NoteCmd)
}
