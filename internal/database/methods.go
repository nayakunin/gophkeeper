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

func (s Storage) GetBinaryData(userID int64) ([]BinaryData, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetTextData(userID int64) ([]TextData, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetBankCardDetails(userID int64, cardName string) ([]BankCardDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetLoginPasswordPairs(userID int64, serviceName string) ([]LoginPasswordPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, user_id, service_name, login, encrypted_password, description FROM login_password_pairs WHERE user_id = $1 AND service_name = $2`, userID, serviceName)
	if err != nil {
		return nil, err
	}

	var pairs []LoginPasswordPair
	for rows.Next() {
		var pair LoginPasswordPair
		if err = rows.Scan(&pair.ID, &pair.UserID, &pair.ServiceName, &pair.Login, &pair.EncryptedPassword, &pair.Description); err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func (s Storage) AddLoginPasswordPair(userID int64, serviceName, login, encryptedPassword, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO login_password_pairs (user_id, service_name, login, encrypted_password, description) VALUES ($1, $2, $3, $4, $5)`, userID, serviceName, login, encryptedPassword, description)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) AddBankCardDetails(userID int64, cardName, encryptedCardNumber, encryptedExpiryDate, encryptedCVC, description string) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddBinaryData(userID int64, binary []byte, description string) error {
	//TODO implement me
	panic("implement me")
}

func (s Storage) AddTextData(userID int64, text, description string) error {
	//TODO implement me
	panic("implement me")
}
