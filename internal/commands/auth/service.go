package auth

// LocalStorage is an interface for storing credentials locally
type LocalStorage interface {
	SaveCredentials(token string, encryptionKey []byte) error
	GetCredentials() (string, []byte, error)
	DeleteCredentials() error
}

// Encryption is an interface for encrypting and decrypting data
type Encryption interface {
	GenerateKey() ([]byte, error)
}

// Service is a struct of the grpc
type Service struct {
	storage    LocalStorage
	encryption Encryption
}

// NewService returns a new Service
func NewService(storage LocalStorage, encryption Encryption) Service {
	return Service{
		storage:    storage,
		encryption: encryption,
	}
}
