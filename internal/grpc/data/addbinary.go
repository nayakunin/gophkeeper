package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddBinaryData adds binary data.
func (s *Service) AddBinaryData(ctx context.Context, in *api.AddBinaryDataRequest) (*api.Empty, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	encryptedData, err := s.encryption.Encrypt(in.GetEncryptedData(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt data: %v", err)
	}

	err = s.storage.AddBinaryData(userID, encryptedData, in.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add binary data: %v", err)
	}

	return &api.Empty{}, nil
}
