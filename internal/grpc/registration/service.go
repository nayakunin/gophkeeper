package registration

import (
	api "github.com/nayakunin/gophkeeper/proto"
)

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
}

// NewService returns a new Service.
func NewService(storage Storage, encryption Encryption) *Service {
	return &Service{
		storage:    storage,
		encryption: encryption,
	}
}
