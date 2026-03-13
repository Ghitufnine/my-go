package jwt

import (
	"os"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

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

func ParseToken(tokenStr string) (string, error) {

	token, err := jwtlib.Parse(
		tokenStr,
		func(token *jwtlib.Token) (interface{}, error) {
			return secret, nil
		},
	)

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwtlib.MapClaims); ok {

		userID := claims["user_id"].(string)

		return userID, nil
	}

	return "", err
}
