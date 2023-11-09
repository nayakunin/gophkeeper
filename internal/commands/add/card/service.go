//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package card

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	AddCardData(ctx context.Context, in *generated.AddBankCardDetailRequest) error
	SetToken(token string)
}

// ApiPreparer is an interface for preparing API requests.
type ApiPreparer interface {
	PrepareCardRequest(data *input.ParseCardResult, encryptionKey []byte) (*generated.AddBankCardDetailRequest, error)
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
