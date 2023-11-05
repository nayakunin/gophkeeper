//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	AddTextData(ctx context.Context, in *generated.AddTextDataRequest) error
}

type ApiPreparer interface {
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
	apiPreparer        ApiPreparer
	api                Api
}

func NewService(credentialsService CredentialsService, encryption Encryption, a Api) *Service {
	apiPreparer := api.NewService(encryption)

	return &Service{
		credentialsService: credentialsService,
		encryption:         encryption,
		apiPreparer:        apiPreparer,
		api:                a,
	}
}
