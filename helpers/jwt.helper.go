package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yash-gkmit/NOTE-TAKER/config"
)

func GenerateJWT(userId, email string) (string, error) {

	claims := &jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.GetInstance().JWTSecret
	return token.SignedString([]byte(secretKey))
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(config.GetInstance().JWTSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("unable to parse claims")
	}

	return claims, nil
}
