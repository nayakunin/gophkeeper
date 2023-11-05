//go:generate mockgen -source=api.go -destination=mocks/service.go -package=mocks
package api

import (
	"fmt"
	"os"

	"github.com/nayakunin/gophkeeper/internal/commands/add/binary/input"
	generated "github.com/nayakunin/gophkeeper/proto"
)

// Encryption is an interface for encrypting and decrypting data.
type Encryption interface {
	Encrypt(text, key []byte) ([]byte, error)
}

type Service struct {
	encryption Encryption
}

func NewService(encryption Encryption) *Service {
	return &Service{
		encryption: encryption,
	}
}

var osReadFile = os.ReadFile

func (a *Service) PrepareBinaryRequest(result *input.ParseBinaryResult, encryptionKey []byte) (*generated.AddBinaryDataRequest, error) {
	file, err := osReadFile(result.Filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	encryptedFile, err := a.encryption.Encrypt(file, encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt file: %w", err)
	}

	return &generated.AddBinaryDataRequest{
		EncryptedData: encryptedFile,
		Description:   result.Description,
	}, nil
}
