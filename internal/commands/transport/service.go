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

// Service is a struct of the grpc.
type Service struct {
	token string
}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}

// SetToken sets the token.
func (s *Service) SetToken(token string) {
	s.token = token
}

// AddBinaryData adds binary data.
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

// AddCardData adds card data.
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

// AddPasswordData adds password data.
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

// AddTextData adds text data.
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

// AuthenticateUser authenticates a user.
func (s *Service) AuthenticateUser(ctx context.Context, in *generated.AuthenticateUserRequest) (*generated.AuthenticateUserResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewAuthServiceClient(conn)
	response, err := client.AuthenticateUser(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("could not authenticate user: %w", err)
	}

	return response, nil
}

// RegisterUser registers a user.
func (s *Service) RegisterUser(ctx context.Context, in *generated.RegisterUserRequest) (*generated.RegisterUserResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewRegistrationServiceClient(conn)
	response, err := client.RegisterUser(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("could not register user: %w", err)
	}

	return response, nil
}

// GetBinaryData gets binary data.
func (s *Service) GetBinaryData(ctx context.Context) (*generated.GetBinaryDataResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	response, err := client.GetBinaryData(ctx, &generated.Empty{})
	if err != nil {
		return nil, fmt.Errorf("could not get binary data: %w", err)
	}

	return response, nil
}

// GetCardDetails gets card details.
func (s *Service) GetCardDetails(ctx context.Context, in *generated.GetBankCardDetailsRequest) (*generated.GetBankCardDetailsResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	response, err := client.GetBankCardDetails(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("could not get card data: %w", err)
	}

	return response, nil
}

// GetLoginPasswordPairs gets login password pairs.
func (s *Service) GetLoginPasswordPairs(ctx context.Context, in *generated.GetLoginPasswordPairsRequest) (*generated.GetLoginPasswordPairsResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	response, err := client.GetLoginPasswordPairs(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("could not get password data: %w", err)
	}

	return response, nil
}

// GetTextData gets text data.
func (s *Service) GetTextData(ctx context.Context) (*generated.GetTextDataResponse, error) {
	conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	defer conn.Close()

	client := generated.NewDataServiceClient(conn)
	md := utils.GetRequestMetadata(s.token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	response, err := client.GetTextData(ctx, &generated.Empty{})
	if err != nil {
		return nil, fmt.Errorf("could not get text data: %w", err)
	}

	return response, nil
}
