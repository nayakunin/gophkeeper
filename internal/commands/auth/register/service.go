//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package register

import (
	"context"

	generated "github.com/nayakunin/gophkeeper/proto"
)

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	GenerateKey() ([]byte, error)
}

// Api is an interface for interacting with the API.
type Api interface {
	RegisterUser(ctx context.Context, in *generated.RegisterUserRequest) (*generated.RegisterUserResponse, error)
}

// Storage is an interface for storing credentials.
type Storage interface {
	SaveCredentials(token string, encryptionKey []byte) error
}

// Service is an interface for interacting with the API.
type Service struct {
	api        Api
	storage    Storage
	encryption Encryption
}

// NewService creates a new instance of Service.
func NewService(encryption Encryption, storage Storage, api Api) *Service {
	return &Service{
		encryption: encryption,
		api:        api,
		storage:    storage,
	}
}
