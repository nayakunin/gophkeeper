//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package registration

import (
	api "github.com/nayakunin/gophkeeper/proto"
)

// AuthService is an interface for authentication.
type AuthService interface {
	GenerateJWT(userID int64) (string, error)
	HashPassword(password string) ([]byte, error)
}

// Storage is an interface for storing credentials.
type Storage interface {
	CreateUser(username string, passwordHash, encryptedMasterKey []byte) (int64, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedRegistrationServiceServer
	storage    Storage
	encryption Encryption
	auth       AuthService
}

// NewService returns a new Service.
func NewService(storage Storage, encryption Encryption, a AuthService) *Service {
	return &Service{
		storage:    storage,
		encryption: encryption,
		auth:       a,
	}
}
