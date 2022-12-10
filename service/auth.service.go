package service

import (
	"context"
	"math/rand"
	"os"

	"time"

	"github.com/go-graphql/auth"
	"github.com/go-graphql/errors"
	"github.com/go-graphql/models"
	"github.com/golang-jwt/jwt"
)

// AuthServiceInterface - .
type AuthServiceInterface interface {
	Login(ctx context.Context, input models.LoginInput) (string, error)
}

// AuthService -
type AuthService struct {
	Issuer    string
	SecretKey string
}

// NewAuthService - .
func NewAuthService() AuthServiceInterface {
	return &AuthService{
		Issuer:    os.Getenv("JWT_ISSUER"),
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

// Login - login user
func (a *AuthService) Login(ctx context.Context, input models.LoginInput) (string, error) {
	if input.Username != os.Getenv("LOGIN_USERNAME") || input.Password != os.Getenv("LOGIN_PASSWORD") {
		return "", errors.ErrorAuth
	}

	user := auth.User{
		UserID:   rand.Intn(100),
		Name:     "Felix",
		Username: input.Username,
		IsAdmin:  true,
	}

	token, err := a.GenerateUserToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateUserToken - generate user jwt token.
func (a *AuthService) GenerateUserToken(user auth.User) (signedToken string, err error) {
	expiredAt := time.Now().Local().Add(time.Hour * 1).Unix()
	claims := &auth.JwtClaims{
		UserID:    user.UserID,
		Username:  user.Username,
		Name:      user.Name,
		IsAdmin:   user.IsAdmin,
		ExpiresAt: expiredAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    a.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
