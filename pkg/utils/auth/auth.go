package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nayakunin/gophkeeper/constants"
	"golang.org/x/crypto/bcrypt"
)

type tokenInfoKey struct {
	name string
}

// UserIDKey is the key for the user id in the context.
var UserIDKey = &tokenInfoKey{"userID"}

// CustomClaims is a struct for JWT token claims.
type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates JWT token.
func GenerateJWT(userID int64) (string, error) {
	claims := CustomClaims{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SecretKey))
}

// ComparePassword compares hash and password.
func ComparePassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

// HashPassword hashes password.
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// ParseToken validates and parses the JWT token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// UserClaimFromToken extracts user identifier from token claims
func UserClaimFromToken(claims *CustomClaims) int64 {
	return claims.UserID
}
