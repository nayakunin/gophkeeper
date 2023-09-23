package add

import "github.com/spf13/cobra"

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
}

func init() {
	AddCmd.AddCommand(PasswordCmd)
	AddCmd.AddCommand(CardCmd)
	AddCmd.AddCommand(BinaryCmd)
	AddCmd.AddCommand(NoteCmd)
}
