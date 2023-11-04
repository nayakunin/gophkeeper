package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetTextData returns text data.
func (s *Service) GetTextData(ctx context.Context, _ *api.Empty) (*api.GetTextDataResponse, error) {
	userID, ok := ctx.Value(authcommon.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	textData, err := s.storage.GetTextData(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get text data: %v", err)
	}

	result := make([]*api.GetTextDataResponseItem, len(textData))
	for i, data := range textData {
		clientEncryptedText, err := s.encryption.Decrypt(data.EncryptedText, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt text: %v", err)
		}

		result[i] = &api.GetTextDataResponseItem{
			Id:            int64(data.ID),
			Description:   data.Description,
			EncryptedText: clientEncryptedText,
		}
	}

	return &api.GetTextDataResponse{
		TextData: result,
	}, nil
}
