package database

import "time"

// User represents a user in the system
type User struct {
	ID                 int       `json:"id"`
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	PasswordHash       []byte    `json:"-"`
	EncryptedMasterKey []byte    `json:"-"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// LoginPasswordPair represents a login/password pair stored by a user
type LoginPasswordPair struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	ServiceName       string    `json:"service_name"`
	Login             string    `json:"login"`
	Description       string    `json:"description"`
	EncryptedPassword []byte    `json:"encrypted_password"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// TextData represents arbitrary textual data stored by a user
type TextData struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Description   string    `json:"description"`
	EncryptedText []byte    `json:"encrypted_text"`
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
	EncryptedCardNumber []byte    `json:"encrypted_card_number"`
	EncryptedExpiryDate []byte    `json:"encrypted_expiry_date"`
	EncryptedCVC        []byte    `json:"encrypted_cvc"`
	Description         string    `json:"description"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
