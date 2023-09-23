package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/nayakunin/gophkeeper/internal/middlewares"
	authService "github.com/nayakunin/gophkeeper/internal/services/auth"
	dataService "github.com/nayakunin/gophkeeper/internal/services/data"
	registrationService "github.com/nayakunin/gophkeeper/internal/services/registration"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
)

const (
	component = "gophkeeper"
	grpcAddr  = ":8080"
)

// interceptorLogger adapts go-kit logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = level.Debug(l).Log(largs...)
		case logging.LevelInfo:
			_ = level.Info(l).Log(largs...)
		case logging.LevelWarn:
			_ = level.Warn(l).Log(largs...)
		case logging.LevelError:
			_ = level.Error(l).Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	rpcLogger := log.With(logger, "service", "gRPC/server", "component", component)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	allButRegistration := func(ctx context.Context, callMeta interceptors.CallMeta) bool {
		return api.RegistrationService_ServiceDesc.ServiceName != callMeta.Service
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(interceptorLogger(rpcLogger)),
			selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(middlewares.Auth), selector.MatchFunc(allButRegistration)),
		))

	api.RegisterRegistrationServiceServer(s, registrationService.NewService())
	api.RegisterAuthServiceServer(s, authService.NewService())
	api.RegisterDataServiceServer(s, dataService.NewService())

	level.Info(logger).Log("msg", "starting gRPC server", "addr", grpcAddr)
	if err := s.Serve(lis); err != nil {
		level.Error(logger).Log("msg", "failed to serve", "err", err)
	}
}
