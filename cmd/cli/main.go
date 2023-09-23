package main

import (
	"context"
	"log"

	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Could not close connection: %v", err)
		}
	}(conn)

	userClient := api.NewUserServiceClient(conn)

	response, err := userClient.RegisterUser(context.Background(), &api.RegisterUserRequest{
		Username: "username",
		Email:    "email@example.com",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("Could not register user: %v", err)
	}

	log.Printf("Registration result: %s", response.GetMessage())

	// Implement other service calls similarly
}
