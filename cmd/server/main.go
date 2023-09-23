package main

import (
	"context"
	"log"
	"net"

	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedUserServiceServer
	api.UnimplementedDataServiceServer
}

func (s *server) RegisterUser(ctx context.Context, req *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	// Implement your logic here
	return &api.RegisterUserResponse{Message: "User Registered", Success: true}, nil
}

func (s *server) AuthenticateUser(ctx context.Context, req *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	// Implement your logic here
	return &api.AuthenticateUserResponse{Token: "some_token", Success: true}, nil
}

// Implement other service methods similarly

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, &server{})
	api.RegisterDataServiceServer(s, &server{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
