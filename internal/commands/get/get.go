package get

import (
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an entry",
}

func init() {
	GetCmd.AddCommand(PasswordCmd)
	GetCmd.AddCommand(CardCmd)
	GetCmd.AddCommand(BinaryCmd)
	GetCmd.AddCommand(NoteCmd)
}
