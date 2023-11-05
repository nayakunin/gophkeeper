//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package logout

type Storage interface {
	DeleteCredentials() error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}