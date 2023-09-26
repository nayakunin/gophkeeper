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

func (s *Service) GetTextData(ctx context.Context, in *api.Empty) (*api.GetTextDataResponse, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	textData, err := s.Storage.GetTextData(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get text data: %v", err)
	}

	result := make([]*api.GetTextDataResponseItem, len(textData))
	for i, data := range textData {
		clientEncryptedText, err := encryption.Decrypt(data.EncryptedText, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt text: %v", err)
		}

		result[i] = &api.GetTextDataResponseItem{
			Description:   data.Description,
			EncryptedText: clientEncryptedText,
		}
	}

	return &api.GetTextDataResponse{
		TextData: result,
	}, nil
}
