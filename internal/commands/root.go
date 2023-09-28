package commands

import (
	"fmt"
	"os"

	"github.com/nayakunin/gophkeeper/internal/commands/add"
	"github.com/nayakunin/gophkeeper/internal/commands/auth"
	"github.com/nayakunin/gophkeeper/internal/commands/get"
	"github.com/spf13/cobra"
)

type LocalStorage interface {
	SaveCredentials(token string, encryptionKey []byte) error
	GetCredentials() (string, []byte, error)
	DeleteCredentials() error
}

type Encryption interface {
	GenerateKey() ([]byte, error)
	Encrypt(text, key []byte) ([]byte, error)
	Decrypt(text, key []byte) ([]byte, error)
}

type Root struct {
	cmd          *cobra.Command
	localStorage LocalStorage
}

func NewRoot(localStorage LocalStorage, encryption Encryption) Root {
	addService := add.NewService(localStorage, encryption)
	authService := auth.NewService(localStorage, encryption)
	getService := get.NewService(localStorage, encryption)

	rootCmd := &cobra.Command{
		Use:   "gophkeeper",
		Short: "Gophkeeper is a tool for managing your gophers",
		Long:  `Gophkeeper is a tool for managing your gophers.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	// Root level commands
	rootCmd.AddCommand(authService.RegisterCmd())
	rootCmd.AddCommand(authService.LoginCmd())
	rootCmd.AddCommand(authService.LogoutCmd())

	// Add subcommands
	rootCmd.AddCommand(addService.Handle())

	// Get subcommands
	rootCmd.AddCommand(getService.Handle())

	return Root{
		cmd: rootCmd,
	}
}

func (r Root) Execute() {
	if err := r.cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
