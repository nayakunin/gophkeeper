package registration

import (
	api "github.com/nayakunin/gophkeeper/proto"
)

type Storage interface {
	CreateUser(username, passwordHash, encryptedMasterKey string) (int64, error)
}

type Encryption interface {
	Encrypt(text string, key []byte) (string, error)
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