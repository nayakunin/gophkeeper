package logout

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (s *Service) GetCmd() *cobra.Command {
	logoutCmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout of the CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := s.storage.DeleteCredentials()
			if err != nil {
				return fmt.Errorf("unable to logout: %w", err)
			}

			fmt.Println("Successfully logged out")
			return nil
		},
	}

	return logoutCmd
}
