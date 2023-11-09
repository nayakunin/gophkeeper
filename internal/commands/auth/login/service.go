//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package login

import (
	"context"

	generated "github.com/nayakunin/gophkeeper/proto"
)

// Api is an interface for interacting with the API.
type Api interface {
	AuthenticateUser(ctx context.Context, in *generated.AuthenticateUserRequest) (*generated.AuthenticateUserResponse, error)
}

// Storage is an interface for storing credentials.
type Storage interface {
	SaveCredentials(token string, encryptionKey []byte) error
}

// Service is an interface for interacting with the API.
type Service struct {
	storage Storage
	api     Api
}

// NewService creates a new instance of Service.
func NewService(storage Storage, api Api) *Service {
	return &Service{
		api:     api,
		storage: storage,
	}
}
