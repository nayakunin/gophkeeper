package auth

import (
	"context"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	api "github.com/nayakunin/gophkeeper/proto"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AuthenticateUser(ctx context.Context, in *api.AuthenticateUserRequest) (*api.AuthenticateUserResponse, error) {
	baseResponse := &api.AuthenticateUserResponse{}

	user, err := s.Storage.GetUser(in.Username)
	if err != nil {
		return baseResponse, err
	}

	if err := ComparePassword(in.Password, user.PasswordHash); err != nil {
		return baseResponse, err
	}

	jwt, err := generateJWT(strconv.Itoa(user.ID))
	if err != nil {
		return baseResponse, err
	}

	return &api.AuthenticateUserResponse{
		Token:   jwt,
		Success: true,
	}, nil
}

func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

type MyCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func generateJWT(userID string) (string, error) {
	claims := MyCustomClaims{
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "secret"

	return token.SignedString([]byte(secretKey))
}
