package auth

type LocalStorage interface {
	SaveCredentials(token string, encryptionKey []byte) error
	GetCredentials() (string, []byte, error)
	DeleteCredentials() error
}

type Service struct {
	storage LocalStorage
}

func NewService(storage LocalStorage) Service {
	return Service{
		storage: storage,
	}
}
