package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/internal/services/auth"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddTextData adds text data.
func (s *Service) AddTextData(ctx context.Context, in *api.AddTextDataRequest) (*api.Empty, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	encryptedText, err := s.encryption.Encrypt(in.GetEncryptedText(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt text: %v", err)
	}

	err = s.storage.AddTextData(userID, encryptedText, in.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add text data: %v", err)
	}

	return &api.Empty{}, nil
}
