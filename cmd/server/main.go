package main

import (
	"log"
	"net"

	"github.com/nayakunin/gophkeeper/internal/services/data"
	"github.com/nayakunin/gophkeeper/internal/services/user"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
)

// Implement other service methods similarly

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, user.NewService())
	api.RegisterDataServiceServer(s, data.NewService())

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
