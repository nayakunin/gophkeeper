package middlewares

import "github.com/nayakunin/gophkeeper/pkg/utils/authcommon"

type AuthClient interface {
	ParseToken(token string) (*authcommon.CustomClaims, error)
	UserClaimFromToken(claims *authcommon.CustomClaims) int64
}

type Service struct {
	a AuthClient
}

func NewService(a AuthClient) *Service {
	return &Service{a: a}
}
