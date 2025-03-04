package util

import (
	"github.com/dgrijalva/jwt-go"
	"rapicreds-backend/src/app/domain"
	"time"
)

var JwtKey = []byte("your_secret_key")

func GenerateJWT(user domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   string(user.ID),
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
