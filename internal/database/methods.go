package database

import (
	"context"

	api "github.com/nayakunin/gophkeeper/proto"
)

func (s Storage) CreateUser(user *api.RegisterUserRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO users (username, password) VALUES ($2, $3)`, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
