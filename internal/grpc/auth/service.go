package auth

import (
	"github.com/nayakunin/gophkeeper/internal/database"
	api "github.com/nayakunin/gophkeeper/proto"
)

// Storage is an interface for storing credentials.
type Storage interface {
	GetUser(username string) (*database.User, error)
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

type AuthService interface {
	ComparePassword(hash, password []byte) error
	GenerateJWT(userID int64) (string, error)
}

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedAuthServiceServer
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
