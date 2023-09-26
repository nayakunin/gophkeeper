package data

import (
	"github.com/nayakunin/gophkeeper/internal/database"
	api "github.com/nayakunin/gophkeeper/proto"
)

type Storage interface {
	GetBinaryData(userID int64) ([]database.BinaryData, error)
	GetTextData(userID int64) ([]database.TextData, error)
	GetBankCardDetails(userID int64, cardName string) ([]database.BankCardDetail, error)
	GetLoginPasswordPairs(userID int64, serviceName string) ([]database.LoginPasswordPair, error)
	AddLoginPasswordPair(userID int64, serviceName, login, encryptedPassword, description string) error
	AddBankCardDetails(userID int64, cardName, encryptedCardNumber, encryptedExpiryDate, encryptedCVC, description string) error
	AddBinaryData(userID int64, binary []byte, description string) error
	AddTextData(userID int64, text, description string) error
}

// Service TODO: Add encryption as a service
// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedDataServiceServer
	Storage Storage
}

// NewService returns a new Service.
func NewService(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
