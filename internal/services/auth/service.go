package auth

import (
	"github.com/nayakunin/gophkeeper/internal/database"
	api "github.com/nayakunin/gophkeeper/proto"
)

type Storage interface {
	GetUser(username string) (*database.User, error)
}

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedAuthServiceServer
	Storage Storage
}

// NewService returns a new Service.
func NewService(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
