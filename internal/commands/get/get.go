package get

import (
	"github.com/spf13/cobra"
)

type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	credentialsService CredentialsService
	encryption         Encryption
}

func NewService(credentialsService CredentialsService, encryption Encryption) *Service {
	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
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
