package transport

import (
	"context"
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	generated "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	token string
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetToken(token string) {
	s.token = token
}

func (s *Service) AddBinaryData(ctx context.Context, in *generated.AddBinaryDataRequest) error {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = client.AddBinaryData(ctx, in)
	if err != nil {
		return fmt.Errorf("could not add binary data: %w", err)
	}

	return nil
}

func (s *Service) AddCardData(ctx context.Context, in *generated.AddBankCardDetailRequest) error {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = client.AddBankCardDetail(ctx, in)
	if err != nil {
		return fmt.Errorf("could not add card data: %w", err)
	}

	return nil
}

func (s *Service) AddPasswordData(ctx context.Context, in *generated.AddLoginPasswordPairRequest) error {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = client.AddLoginPasswordPair(ctx, in)
	if err != nil {
		return fmt.Errorf("could not add password data: %w", err)
	}

	return nil
}

func (s *Service) AddTextData(ctx context.Context, in *generated.AddTextDataRequest) error {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	_, err = client.AddTextData(ctx, in)
	if err != nil {
		return fmt.Errorf("could not add text data: %w", err)
	}

	return nil
}