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

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedAuthServiceServer
	storage    Storage
	encryption Encryption
}

// NewService returns a new Service.
func NewService(storage Storage, encryption Encryption) *Service {
	return &Service{
		storage:    storage,
		encryption: encryption,
	}
}
