//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package card

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/card/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	GetCardDetails(ctx context.Context, in *generated.GetBankCardDetailsRequest) (*generated.GetBankCardDetailsResponse, error)
}

type Output interface {
	MakeResponse(response *generated.GetBankCardDetailsResponse, encryptionKey []byte) ([]output.CardResult, error)
}

type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	output             Output
	credentialsService CredentialsService
	encryption         Encryption
	api                Api
}

func NewService(encryption Encryption, credentialsService CredentialsService, api Api) *Service {
	out := output.NewService(encryption)

	return &Service{
		output:             out,
		credentialsService: credentialsService,
		encryption:         encryption,
		api:                api,
	}
}
