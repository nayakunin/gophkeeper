//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package register

import (
	"context"

	generated "github.com/nayakunin/gophkeeper/proto"
)

type Encryption interface {
	GenerateKey() ([]byte, error)
}

type Api interface {
	RegisterUser(ctx context.Context, in *generated.RegisterUserRequest) (*generated.RegisterUserResponse, error)
}

type Storage interface {
	SaveCredentials(token string, encryptionKey []byte) error
}

type Service struct {
	api        Api
	storage    Storage
	encryption Encryption
}

func NewService(encryption Encryption, storage Storage, api Api) *Service {
	return &Service{
		encryption: encryption,
		api:        api,
		storage:    storage,
	}
}
