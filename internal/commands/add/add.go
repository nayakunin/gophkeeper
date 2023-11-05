//go:generate mockgen -source=add.go -destination=mocks/service.go -package=mocks
package add

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary"
	"github.com/nayakunin/gophkeeper/internal/commands/add/card"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

type Api interface {
	AddBinaryData(ctx context.Context, in *generated.AddBinaryDataRequest) error
	AddCardData(ctx context.Context, in *generated.AddBankCardDetailRequest) error
	AddPasswordData(ctx context.Context, in *generated.AddLoginPasswordPairRequest) error
}

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
	api                Api
}

// NewService returns a new Service.
func NewService(credentialsService CredentialsService, encryption Encryption, api Api) *Service {
	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
		api:                api,
	}
}

// Handle returns a new cobra command.
func (s *Service) Handle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new entry",
	}

	binaryService := binary.NewService(s.credentialsService, s.encryption, s.api)
	cardService := card.NewService(s.credentialsService, s.encryption, s.api)
	passwordService := password.NewService(s.credentialsService, s.encryption, s.api)
	textService := text.NewService(s.credentialsService, s.encryption)

	cmd.AddCommand(passwordService.GetCmd())
	cmd.AddCommand(binaryService.GetCmd())
	cmd.AddCommand(textService.GetCmd())
	cmd.AddCommand(cardService.GetCmd())

	return cmd
}
