package authcommon

import (
	"github.com/golang-jwt/jwt/v5"
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
