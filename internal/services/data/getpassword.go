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

func (s *Service) GetLoginPasswordPairs(ctx context.Context, in *api.GetLoginPasswordPairsRequest) (*api.GetLoginPasswordPairsResponse, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	pairs, err := s.Storage.GetLoginPasswordPairs(userID, in.GetServiceName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get login password pairs: %v", err)
	}

	result := make([]*api.LoginPasswordPair, len(pairs))
	for i, pair := range pairs {
		clientEncryptedPassword, err := encryption.Decrypt(pair.EncryptedPassword, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt password: %v", err)
		}

		result[i] = &api.LoginPasswordPair{
			Id:                int64(pair.ID),
			Login:             pair.Login,
			EncryptedPassword: clientEncryptedPassword,
			Description:       pair.Description,
			ServiceName:       pair.ServiceName,
		}
	}

	return &api.GetLoginPasswordPairsResponse{
		LoginPasswordPairs: result,
	}, nil
}
