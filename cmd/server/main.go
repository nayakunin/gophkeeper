package main

import (
	"context"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/database"
	authGrpcService "github.com/nayakunin/gophkeeper/internal/grpc/auth"
	dataGrpcService "github.com/nayakunin/gophkeeper/internal/grpc/data"
	registrationGrpcService "github.com/nayakunin/gophkeeper/internal/grpc/registration"
	"github.com/nayakunin/gophkeeper/internal/middlewares"
	"github.com/nayakunin/gophkeeper/internal/services/auth"
	"github.com/nayakunin/gophkeeper/internal/services/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", constants.GrpcPort))
	if err != nil {
		fmt.Println("failed to listen", "err", err)
		panic(err)
	}

	allButAuth := func(ctx context.Context, callMeta interceptors.CallMeta) bool {
		return api.RegistrationService_ServiceDesc.ServiceName != callMeta.Service && api.AuthService_ServiceDesc.ServiceName != callMeta.Service
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(grpcAuth.UnaryServerInterceptor(middlewares.Auth), selector.MatchFunc(allButAuth)),
		))

	ctx, cancel := context.WithTimeout(context.Background(), constants.DBTimeout)
	defer cancel()
	storage, err := database.NewStorage(ctx, "postgresql://localhost:5432/postgres")
	if err != nil {
		fmt.Println("failed to connect to database", "err", err)
		panic(err)
	}

	encryptionService := encryption.NewService()
	authService := auth.NewService()

	api.RegisterRegistrationServiceServer(s, registrationGrpcService.NewService(storage, encryptionService, authService))
	api.RegisterAuthServiceServer(s, authGrpcService.NewService(storage, encryptionService, authService))
	api.RegisterDataServiceServer(s, dataGrpcService.NewService(storage, encryptionService))

	fmt.Println("starting gRPC server", "addr", constants.GrpcPort)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve", "err", err)
	}
}
