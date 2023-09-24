package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (s *Service) textCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "text",
		Short: "Add a new text data",
		Run: func(cmd *cobra.Command, args []string) {
			// Logic for adding a password
			fmt.Println("Adding text data...")
		},
	}
}
