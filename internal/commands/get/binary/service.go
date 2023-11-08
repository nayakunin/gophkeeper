//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package binary

import (
	"context"

	"github.com/nayakunin/gophkeeper/internal/commands/get/binary/output"
	generated "github.com/nayakunin/gophkeeper/proto"
)

type Api interface {
	GetBinaryData(ctx context.Context) (*generated.GetBinaryDataResponse, error)
}

type Output interface {
	MakeResponse(response *generated.GetBinaryDataResponse, encryptionKey []byte) ([]output.BinaryResult, error)
}

type CredentialsService interface {
	GetCredentials() (string, []byte, error)
}

type Encryption interface {
	Decrypt(data []byte, key []byte) ([]byte, error)
}

type Service struct {
	encryption         Encryption
	credentialsService CredentialsService
	output             Output
	api                Api
}

func NewService(encryption Encryption, credentialsService CredentialsService, api Api) *Service {
	op := output.NewService(encryption)

	return &Service{
		encryption:         encryption,
		credentialsService: credentialsService,
		output:             op,
		api:                api,
	}
}
