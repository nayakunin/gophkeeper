//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package binary

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	GetBinaryData(ctx context.Context) (*generated.GetBinaryDataResponse, error)
	SetToken(token string)
}

// Output is an interface for preparing output.
type Output interface {
	MakeResponse(response *generated.GetBinaryDataResponse, encryptionKey []byte) ([]output.BinaryResult, error)
}

// CredentialsService is an interface for getting credentials.
type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(data []byte, key []byte) ([]byte, error)
}

// Service is an interface for interacting with the API.
type Service struct {
	encryption         Encryption
	credentialsService CredentialsService
	output             Output
	api                Api
}

// NewService creates a new instance of Service.
func NewService(encryption Encryption, credentialsService CredentialsService, api Api) *Service {
	op := output.NewService(encryption)

	return &Service{
		encryption:         encryption,
		credentialsService: credentialsService,
		output:             op,
		api:                api,
	}
}
