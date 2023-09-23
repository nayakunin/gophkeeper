package registration

import (
	"context"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s *Service) RegisterUser(ctx context.Context, in *api.RegisterUserRequest) (*api.RegisterUserResponse, error) {
	// ...
	return nil, nil
}
