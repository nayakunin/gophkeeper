//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package card

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/card/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	GetCardDetails(ctx context.Context, in *generated.GetBankCardDetailsRequest) (*generated.GetBankCardDetailsResponse, error)
	SetToken(token string)
}

// Output is an interface for preparing output.
type Output interface {
	MakeResponse(response *generated.GetBankCardDetailsResponse, encryptionKey []byte) ([]output.CardResult, error)
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
