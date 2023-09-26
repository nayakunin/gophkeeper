package auth

type LocalStorage interface {
	SaveCredentials(token string, encryptionKey []byte) error
	GetCredentials() (string, []byte, error)
	DeleteCredentials() error
}

type Encryption interface {
	GenerateKey() ([]byte, error)
}

type Service struct {
	storage    LocalStorage
	encryption Encryption
}

func NewService(storage LocalStorage, encryption Encryption) Service {
	return Service{
		storage:    storage,
		encryption: encryption,
	}
}
