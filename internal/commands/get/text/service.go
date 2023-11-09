//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/text/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	GetTextData(ctx context.Context) (*generated.GetTextDataResponse, error)
	SetToken(token string)
}

// Output is an interface for preparing output.
type Output interface {
	MakeResponse(response *generated.GetTextDataResponse, encryptionKey []byte) ([]output.TextResult, error)
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
	api                Api
	output             Output
}

// NewService returns a new Service.
func NewService(encryption Encryption, credentialsService CredentialsService, api Api) *Service {
	out := output.NewService(encryption)

	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
		api:                api,
		output:             out,
	}
}
