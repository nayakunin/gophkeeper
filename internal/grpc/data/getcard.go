package data

import (
	"context"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/auth"
	api "github.com/nayakunin/gophkeeper/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetBankCardDetails returns bank card details.
func (s *Service) GetBankCardDetails(ctx context.Context, in *api.GetBankCardDetailsRequest) (*api.GetBankCardDetailsResponse, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return nil, status.Errorf(codes.Internal, "userID not found in context")
	}

	cardDetails, err := s.storage.GetBankCardDetails(userID, in.GetCardName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get bank card details: %v", err)
	}

	result := make([]*api.BankCardDetail, len(cardDetails))
	for i, card := range cardDetails {
		clientEncryptedCardNumber, err := s.encryption.Decrypt(card.EncryptedCardNumber, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt card number: %v", err)
		}

		clientEncryptedExpiryDate, err := s.encryption.Decrypt(card.EncryptedExpiryDate, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt card expiration: %v", err)
		}

		clientEncryptedCVC, err := s.encryption.Decrypt(card.EncryptedCVC, []byte(constants.EncryptionKey))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decrypt card cvc: %v", err)
		}

		result[i] = &api.BankCardDetail{
			CardName:            card.CardName,
			EncryptedCardNumber: clientEncryptedCardNumber,
			EncryptedExpiryDate: clientEncryptedExpiryDate,
			EncryptedCvc:        clientEncryptedCVC,
			Description:         card.Description,
		}
	}

	return &api.GetBankCardDetailsResponse{
		BankCardDetails: result,
	}, nil
}
