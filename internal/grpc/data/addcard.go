package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AddBankCardDetail(ctx context.Context, in *api.AddBankCardDetailRequest) (*api.Empty, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	encryptedCardNumber, err := s.encryption.Encrypt(in.GetEncryptedCardNumber(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt card number: %v", err)
	}

	encryptedExpiration, err := s.encryption.Encrypt(in.GetEncryptedExpiryDate(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt card expiration: %v", err)
	}

	encryptedCVC, err := s.encryption.Encrypt(in.GetEncryptedCvc(), []byte(constants.EncryptionKey))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to encrypt card cvc: %v", err)
	}

	err = s.storage.AddBankCardDetails(userID, in.GetCardName(), encryptedCardNumber, encryptedExpiration, encryptedCVC, in.GetDescription())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add bank card detail: %v", err)
	}

	return &api.Empty{}, nil
}
