package utils

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string, timeToExpire int64) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      timeToExpire,
	}

	// Generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
