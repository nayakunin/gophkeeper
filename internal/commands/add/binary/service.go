//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package binary

import (
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	"github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	PrepareBinaryRequest(result *input.ParseBinaryResult, encryptionKey []byte) (*generated.AddBinaryDataRequest, error)
}

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	credentialsService CredentialsService
	encryption         Encryption
	apiPreparer        Api
}

func NewService(credentialsService CredentialsService, encryption Encryption) *Service {
	apiPreparer := api.NewService(encryption)

	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
		apiPreparer:        apiPreparer,
	}
}
