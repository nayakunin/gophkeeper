package auth

import api "github.com/nayakunin/gophkeeper/proto"

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedAuthServiceServer
}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}
