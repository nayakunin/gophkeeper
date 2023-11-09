//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	AddTextData(ctx context.Context, in *generated.AddTextDataRequest) error
	SetToken(token string)
}

// ApiPreparer is an interface for preparing API requests.
type ApiPreparer interface {
	PrepareTextRequest(result *input.ParseTextResult, encryptionKey []byte) (*generated.AddTextDataRequest, error)
}

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

// Service is an interface for interacting with the API.
type Service struct {
	credentialsService CredentialsService
	encryption         Encryption
	apiPreparer        ApiPreparer
	api                Api
}

// NewService creates a new instance of Service.
func NewService(credentialsService CredentialsService, encryption Encryption, a Api) *Service {
	apiPreparer := api.NewService(encryption)

	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
		apiPreparer:        apiPreparer,
		api:                a,
	}
}
