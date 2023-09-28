package add

import "github.com/spf13/cobra"

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

// Service is a struct of the grpc.
type Service struct {
	credentialsService CredentialsService
	encryption         Encryption
}

// NewService returns a new Service.
func NewService(credentialsService CredentialsService, encryption Encryption) *Service {
	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
	}
}

// Handle returns a new cobra command.
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
