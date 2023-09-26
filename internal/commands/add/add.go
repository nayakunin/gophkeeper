package add

import "github.com/spf13/cobra"

type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

type Encryption interface {
	Encrypt(text string, key []byte) (string, error)
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
		Use:   "add",
		Short: "Add a new entry",
	}

	cmd.AddCommand(s.passwordCmd())
	cmd.AddCommand(s.binaryCmd())
	cmd.AddCommand(s.textCmd())
	cmd.AddCommand(s.cardCmd())

	return cmd
}
