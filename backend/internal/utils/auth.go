package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string, timeToExpire int64) (string, error) {
	jwtSecretKey := os.Getenv("JWT_KEY")
	if jwtSecretKey == "" {
		return "", errors.New("failed to get jwt secret key")
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      timeToExpire,
	}

	// Generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
