//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	GetLoginPasswordPairs(ctx context.Context, in *generated.GetLoginPasswordPairsRequest) (*generated.GetLoginPasswordPairsResponse, error)
	SetToken(token string)
}

// Output is an interface for preparing output.
type Output interface {
	MakeResponse(response *generated.GetLoginPasswordPairsResponse, encryptionKey []byte) ([]output.PasswordResult, error)
}

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

// Service is an interface for interacting with the API.
type Service struct {
	output             Output
	credentialsService CredentialsService
	encryption         Encryption
	api                Api
}

// NewService creates a new instance of Service.
func NewService(encryption Encryption, credentialsService CredentialsService, api Api) *Service {
	out := output.NewService(encryption)

	return &Service{
		output:             out,
		credentialsService: credentialsService,
		encryption:         encryption,
		api:                api,
	}
}