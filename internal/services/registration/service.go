package registration

import (
	api "github.com/nayakunin/gophkeeper/proto"
)

type Storage interface {
	CreateUser(username, passwordHash, encryptedMasterKey string) (int64, error)
}

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedRegistrationServiceServer
	Storage Storage
}

// NewService returns a new Service.
func NewService(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
