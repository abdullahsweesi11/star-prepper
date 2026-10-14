package utils

import (
	"os"

	jwt5 "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId string, expTime int64) (string, error) {
	generateToken := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"id":  userId,
		"exp": expTime,
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
