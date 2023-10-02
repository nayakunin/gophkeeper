package data

import (
	"github.com/nayakunin/gophkeeper/internal/database"
	api "github.com/nayakunin/gophkeeper/proto"
)

// Storage is an interface for storing credentials.
type Storage interface {
	GetBinaryData(userID int64) ([]database.BinaryData, error)
	GetTextData(userID int64) ([]database.TextData, error)
	GetBankCardDetails(userID int64, cardName string) ([]database.BankCardDetail, error)
	GetLoginPasswordPairs(userID int64, serviceName string) ([]database.LoginPasswordPair, error)
	AddLoginPasswordPair(userID int64, serviceName, login string, encryptedPassword []byte, description string) error
	AddBankCardDetails(userID int64, cardName string, encryptedCardNumber, encryptedExpiryDate, encryptedCVC []byte, description string) error
	AddBinaryData(userID int64, binary []byte, description string) error
	AddTextData(userID int64, text []byte, description string) error
}

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
	Decrypt(text, key []byte) ([]byte, error)
}

type AuthService interface {
}

// Service is a struct of the grpc.
type Service struct {
	api.UnimplementedDataServiceServer
	storage    Storage
	encryption Encryption
}

// NewService returns a new Service.
func NewService(storage Storage, encryption Encryption) *Service {
	return &Service{
		storage:    storage,
		encryption: encryption,
	}
}
