package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (s *Service) cardCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "card",
		Short: "Get a card",
		Run: func(cmd *cobra.Command, args []string) {
			// Logic for getting a password
			fmt.Println("Getting card...")
		},
	}
}
