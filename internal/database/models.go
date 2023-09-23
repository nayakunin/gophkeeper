package database

import "time"

// User represents a user in the system
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// LoginPasswordPair represents a login/password pair stored by a user
type LoginPasswordPair struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	ServiceName  string    `json:"service_name"`
	Login        string    `json:"login"`
	Description  string    `json:"description"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TextData represents arbitrary textual data stored by a user
type TextData struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Description   string    `json:"description"`
	EncryptedText string    `json:"encrypted_text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// BinaryData represents arbitrary binary data stored by a user
type BinaryData struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Description   string    `json:"description"`
	EncryptedData []byte    `json:"encrypted_data"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// BankCardDetail represents bank card details stored by a user
type BankCardDetail struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	CardName            string    `json:"card_name"`
	EncryptedCardNumber string    `json:"encrypted_card_number"`
	EncryptedExpiryDate string    `json:"encrypted_expiry_date"`
	EncryptedCVC        string    `json:"encrypted_cvc"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
