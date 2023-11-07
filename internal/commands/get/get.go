//go:generate mockgen -source=get.go -destination=mocks/service.go -package=mocks
package get

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary"
	"github.com/nayakunin/gophkeeper/internal/commands/get/card"
	"github.com/nayakunin/gophkeeper/internal/commands/get/password"
	"github.com/nayakunin/gophkeeper/internal/commands/get/text"
	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

type Api interface {
	GetBinaryData(ctx context.Context) (*generated.GetBinaryDataResponse, error)
	GetCardDetails(ctx context.Context, in *generated.GetBankCardDetailsRequest) (*generated.GetBankCardDetailsResponse, error)
	GetLoginPasswordPairs(ctx context.Context, in *generated.GetLoginPasswordPairsRequest) (*generated.GetLoginPasswordPairsResponse, error)
	GetTextData(ctx context.Context) (*generated.GetTextDataResponse, error)
}

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

// Service is a struct of the grpc.
type Service struct {
	credentialsService CredentialsService
	encryption         Encryption
	api 							Api
}

// NewService returns a new Service.
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

	binaryService := binary.NewService(s.encryption, s.credentialsService, s.api)
	cardService := card.NewService(s.encryption, s.credentialsService, s.api)
	passwordService := password.NewService(s.encryption, s.credentialsService, s.api)
	textService := text.NewService(s.encryption, s.credentialsService, s.api)

	cmd.AddCommand(binaryService.GetCmd())
	cmd.AddCommand(cardService.GetCmd())
	cmd.AddCommand(passwordService.GetCmd())
	cmd.AddCommand(textService.GetCmd())

	return cmd
}
