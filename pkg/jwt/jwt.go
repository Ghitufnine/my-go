package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var secret = []byte("supersecret")

func GenerateAccessToken(userID string) (string, error) {

	claims := jwtlib.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func GenerateRefreshToken(userID string) (string, error) {

	claims := jwtlib.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
