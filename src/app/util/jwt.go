package util

import (
	"github.com/dgrijalva/jwt-go"
	"rapicreds-backend/src/app/domain"
	"strconv"
	"time"
)

var JwtKey = []byte("your_secret_key")

func GenerateJWT(user domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	subject := strconv.Itoa(int(user.ID))
	println("Subject: ", subject)
	claims := &jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
