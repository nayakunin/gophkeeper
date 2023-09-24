package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	"github.com/nayakunin/gophkeeper/pkg/utils/encryption"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AddLoginPasswordPair(ctx context.Context, in *api.AddLoginPasswordPairRequest) (*api.Empty, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	encryptedPassword, err := encryption.Encrypt(in.GetEncryptedPassword(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	err = s.Storage.AddLoginPasswordPair(userID, in.GetServiceName(), in.GetLogin(), encryptedPassword, in.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add login password pair: %v", err)
	}

	return &api.Empty{}, nil
}
