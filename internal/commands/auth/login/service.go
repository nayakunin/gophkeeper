//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package login

import (
	"context"

	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	AuthenticateUser(ctx context.Context, in *generated.AuthenticateUserRequest) (*generated.AuthenticateUserResponse, error)
}

type Storage interface {
	SaveCredentials(token string, encryptionKey []byte) error
}

type Service struct {
	storage Storage
	api     Api
}

func NewService(storage Storage, api Api) *Service {
	return &Service{
		api:     api,
		storage: storage,
	}
}
