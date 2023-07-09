package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// JwtClaims - .
type JwtClaims struct {
	jwt.StandardClaims
	UserID    int
	Username  string
	Name      string
	IsAdmin   bool
	ExpiresAt int64
}

// ValidateAuthorization - validate jwt token.
func ValidateAuthorization(signedToken string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.StandardClaims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}
