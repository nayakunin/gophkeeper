//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package binary

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	AddBinaryData(ctx context.Context, in *generated.AddBinaryDataRequest) error
}

type ApiPreparer interface {
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
	apiPreparer        ApiPreparer
	api                Api
}

func NewService(c CredentialsService, e Encryption, a Api) *Service {
	ap := api.NewService(e)

	return &Service{
		credentialsService: c,
		encryption:         e,
		apiPreparer:        ap,
		api:                a,
	}
}
