package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetBinaryData returns binary data.
func (s *Service) GetBinaryData(ctx context.Context, in *api.Empty) (*api.GetBinaryDataResponse, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	data, err := s.storage.GetBinaryData(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get binary data: %v", err)
	}

	result := make([]*api.GetBinaryDataResponseItem, len(data))
	for i, pair := range data {
		clientEncryptedData, err := s.encryption.Decrypt(pair.EncryptedData, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt data: %v", err)
		}

		result[i] = &api.GetBinaryDataResponseItem{
			Id:            int64(pair.ID),
			EncryptedData: clientEncryptedData,
			Description:   pair.Description,
		}
	}

	return &api.GetBinaryDataResponse{
		BinaryData: result,
	}, nil
}
