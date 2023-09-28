package database

import (
	"context"
	"fmt"
)

// CreateUser creates a new user in the database
func (s Storage) CreateUser(username string, passwordHash []byte, encryptedMasterKey []byte) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	var userID int64
	err = conn.QueryRow(ctx, `INSERT INTO users (username, password_hash, encrypted_master_key) VALUES ($1, $2, $3) RETURNING id`, username, passwordHash, encryptedMasterKey).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("could not insert user: %w", err)
	}

	return userID, nil
}

// GetUser returns a user from the database
func (s Storage) GetUser(username string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	var user User
	if err = conn.QueryRow(ctx, `SELECT id, username, password_hash, encrypted_master_key FROM users WHERE username = $1`, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.EncryptedMasterKey); err != nil {
		return nil, fmt.Errorf("could not select user: %w", err)
	}

	return &user, nil
}

// GetBinaryData returns binary data from the database
func (s Storage) GetBinaryData(userID int64) ([]BinaryData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, user_id, encrypted_data, description FROM binary_data WHERE user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("could not select binary data: %w", err)
	}

	var binaryData []BinaryData
	for rows.Next() {
		var data BinaryData
		if err = rows.Scan(&data.ID, &data.UserID, &data.EncryptedData, &data.Description); err != nil {
			return nil, fmt.Errorf("could not scan binary data: %w", err)
		}
		binaryData = append(binaryData, data)
	}

	return binaryData, nil
}

// GetTextData returns text data from the database
func (s Storage) GetTextData(userID int64) ([]TextData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, user_id, encrypted_text, description FROM text_data WHERE user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("could not select text data: %w", err)
	}

	var textData []TextData
	for rows.Next() {
		var data TextData
		if err = rows.Scan(&data.ID, &data.UserID, &data.EncryptedText, &data.Description); err != nil {
			return nil, fmt.Errorf("could not scan text data: %w", err)
		}
		textData = append(textData, data)
	}

	return textData, nil
}

// GetBankCardDetails returns bank card details from the database
func (s Storage) GetBankCardDetails(userID int64, cardName string) ([]BankCardDetail, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, user_id, card_name, encrypted_card_number, encrypted_expiry_date, encrypted_cvc, description FROM bank_card_details WHERE user_id = $1 AND card_name = $2`, userID, cardName)
	if err != nil {
		return nil, fmt.Errorf("could not select bank card details: %w", err)
	}

	var details []BankCardDetail
	for rows.Next() {
		var detail BankCardDetail
		if err = rows.Scan(&detail.ID, &detail.UserID, &detail.CardName, &detail.EncryptedCardNumber, &detail.EncryptedExpiryDate, &detail.EncryptedCVC, &detail.Description); err != nil {
			return nil, fmt.Errorf("could not scan bank card detail: %w", err)
		}
		details = append(details, detail)
	}

	return details, nil
}

// GetLoginPasswordPairs returns login password pairs from the database
func (s Storage) GetLoginPasswordPairs(userID int64, serviceName string) ([]LoginPasswordPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `SELECT id, user_id, service_name, login, encrypted_password, description FROM login_password_pairs WHERE user_id = $1 AND service_name = $2`, userID, serviceName)
	if err != nil {
		return nil, fmt.Errorf("could not select login password pairs: %w", err)
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

// AddLoginPasswordPair adds a new login password pair to the database
func (s Storage) AddLoginPasswordPair(userID int64, serviceName, login string, encryptedPassword []byte, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO login_password_pairs (user_id, service_name, login, encrypted_password, description) VALUES ($1, $2, $3, $4, $5)`, userID, serviceName, login, encryptedPassword, description)
	if err != nil {
		return fmt.Errorf("could not insert login password pair: %w", err)
	}

	return nil
}

// AddBankCardDetails adds a new bank card detail to the database
func (s Storage) AddBankCardDetails(userID int64, cardName string, encryptedCardNumber, encryptedExpiryDate, encryptedCVC []byte, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO bank_card_details (user_id, card_name, encrypted_card_number, encrypted_expiry_date, encrypted_cvc, description) VALUES ($1, $2, $3, $4, $5, $6)`, userID, cardName, encryptedCardNumber, encryptedExpiryDate, encryptedCVC, description)
	if err != nil {
		return fmt.Errorf("could not insert bank card detail: %w", err)
	}

	return nil
}

// AddBinaryData adds a new binary data to the database
func (s Storage) AddBinaryData(userID int64, binary []byte, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO binary_data (user_id, encrypted_data, description) VALUES ($1, $2, $3)`, userID, binary, description)
	if err != nil {
		return fmt.Errorf("could not insert binary data: %w", err)
	}

	return nil
}

// AddTextData adds a new text data to the database
func (s Storage) AddTextData(userID int64, text []byte, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	conn, err := s.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("could not acquire connection: %w", err)
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, `INSERT INTO text_data (user_id, encrypted_text, description) VALUES ($1, $2, $3)`, userID, text, description)
	if err != nil {
		return fmt.Errorf("could not insert text data: %w", err)
	}

	return nil
}
