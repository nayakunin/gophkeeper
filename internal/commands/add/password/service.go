//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/add/password/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	AddPasswordData(ctx context.Context, in *generated.AddLoginPasswordPairRequest) error
}

type ApiPreparer interface {
	PreparePasswordRequest(result *input.ParsePasswordResult, encryptionKey []byte) (*generated.AddLoginPasswordPairRequest, error)
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
