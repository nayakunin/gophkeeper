package registration

import (
	api "github.com/nayakunin/gophkeeper/proto"
)

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedRegistrationServiceServer
}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}
