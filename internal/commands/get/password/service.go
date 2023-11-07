//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/password/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	GetLoginPasswordPairs(ctx context.Context, in *generated.GetLoginPasswordPairsRequest) (*generated.GetLoginPasswordPairsResponse, error)
}

type Output interface {
	MakeResponse(response *generated.GetLoginPasswordPairsResponse, encryptionKey []byte) ([]output.PasswordResult, error)
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