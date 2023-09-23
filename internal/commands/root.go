package commands

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/internal/commands/add"
	"github.com/nayakunin/gophkeeper/internal/commands/get"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gophkeeper",
	Short: "Gophkeeper is a tool for managing your gophers",
	Long:  `Gophkeeper is a tool for managing your gophers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(get.GetCmd)
}
