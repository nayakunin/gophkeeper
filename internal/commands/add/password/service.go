//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package password

import (
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/api"
	"github.com/nayakunin/gophkeeper/internal/commands/add/password/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
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
