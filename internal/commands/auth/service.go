package auth

type LocalStorage interface {
	SaveCredentials(token, encryptionKey string) error
	GetCredentials() (string, string, error)
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
