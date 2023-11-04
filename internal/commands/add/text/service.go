//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	PrepareTextRequest(result *input.ParseTextResult, encryptionKey []byte) (*generated.AddTextDataRequest, error)
}

type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

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
