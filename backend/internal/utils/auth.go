package utils

import (
	"errors"
	"myapp/internal/config"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int64, username string, timeToExpire int64) (string, error) {
	claims := jwt.MapClaims{
		"id":       userId,
		"username": username,
		"exp":      timeToExpire,
	}

	// Generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, err
}

func GetUsernameFromParsedToken(parsedToken *jwt.Token) (string, error) {
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims["username"].(string), nil
	}

	return "", errors.New("failed to get username from token")
}

func GetUserIdFromParsedToken(parsedToken *jwt.Token) (int64, error) {
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return int64(claims["id"].(float64)), nil
	}

	return 0, errors.New("failed to get user ID from token")
}
