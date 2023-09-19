package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	secret := os.Getenv("SECRET_KEY")

	token, err := claims.SignedString([]byte(secret))

	return token, err
}
