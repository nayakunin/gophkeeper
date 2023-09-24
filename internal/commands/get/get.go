package get

import (
	"github.com/spf13/cobra"
)

type CredentialsService interface {
	GetCredentials() (string, string, error)
}

type Service struct {
	credentialsService CredentialsService
}

func NewService(credentialsService CredentialsService) *Service {
	return &Service{
		credentialsService: credentialsService,
	}
}

func (s *Service) Handle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an entry",
	}

	cmd.AddCommand(s.passwordCmd())
	cmd.AddCommand(s.binaryCmd())
	cmd.AddCommand(s.textCmd())
	cmd.AddCommand(s.cardCmd())

	return cmd
}
