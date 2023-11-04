//go:generate mockgen -source=add.go -destination=mocks/service.go -package=mocks
package add

import (
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary"
	"github.com/spf13/cobra"
)

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

	binaryService := binary.NewService(s.credentialsService, s.encryption)

	cmd.AddCommand(s.passwordCmd())
	cmd.AddCommand(binaryService.GetBinaryCmd())
	cmd.AddCommand(s.textCmd())
	cmd.AddCommand(s.cardCmd())

	return cmd
}
