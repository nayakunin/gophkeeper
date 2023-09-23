package database

import (
	"context"
)

func (s Storage) CreateUser(username, passwordHash string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO users (username, password_hash) VALUES ($1, $2)`, username, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) GetUser(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var user User
	if err = conn.QueryRow(ctx, `SELECT id, username, password_hash FROM users WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		return nil, err
	}

	return &user, nil
}
