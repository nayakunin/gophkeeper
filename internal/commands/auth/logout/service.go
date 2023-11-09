//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package logout

// Storage is an interface for storing credentials.
type Storage interface {
	DeleteCredentials() error
}

// Service is an interface for interacting with the API.
type Service struct {
	storage Storage
}

// NewService creates a new instance of Service.
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}