package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (s *Service) textCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "text",
		Short: "Get a text data",
		Run: func(cmd *cobra.Command, args []string) {
			// Logic for getting a password
			fmt.Println("Getting text data...")
		},
	}
}
