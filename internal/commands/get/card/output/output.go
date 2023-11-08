//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
)

type Encryption interface {
	Decrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	encryption Encryption
}

func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

type CardResult struct {
	Name        string `json:"label"`
	Number      string `json:"number"`
	Expiration  string `json:"expiration"`
	Cvc         string `json:"cvv"`
	Description string `json:"description"`
}

func (s *Service) MakeResponse(response *generated.GetBankCardDetailsResponse, encryptionKey []byte) ([]CardResult, error) {
	results := make([]CardResult, len(response.GetBankCardDetails()))
	for i, card := range response.GetBankCardDetails() {
		number, err := s.encryption.Decrypt(card.GetEncryptedCardNumber(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card number: %w", err)
		}
		expiration, err := s.encryption.Decrypt(card.GetEncryptedExpiryDate(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card expiration date: %w", err)
		}
		cvc, err := s.encryption.Decrypt(card.GetEncryptedCvc(), encryptionKey)
		if err != nil {
			return nil, fmt.Errorf("could not decrypt card CVC: %w", err)
		}
		results[i] = CardResult{
			Name:        card.GetCardName(),
			Number:      string(number),
			Expiration:  string(expiration),
			Cvc:         string(cvc),
			Description: card.GetDescription(),
		}
	}

	return results, nil
}