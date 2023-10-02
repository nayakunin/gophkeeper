package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GenerateJWT generates JWT token.
func (s *Service) GenerateJWT(userID int64) (string, error) {
	claims := authcommon.CustomClaims{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SecretKey))
}

// ComparePassword compares hash and password.
func (s *Service) ComparePassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

// HashPassword hashes password.
func (s *Service) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// ParseToken validates and parses the JWT token
func (s *Service) ParseToken(tokenStr string) (*authcommon.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &authcommon.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*authcommon.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// UserClaimFromToken extracts user identifier from token claims
func (s *Service) UserClaimFromToken(claims *authcommon.CustomClaims) int64 {
	return claims.UserID
}
