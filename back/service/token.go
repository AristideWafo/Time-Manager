package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	Team string
	jwt.RegisteredClaims
}

var secretKey = []byte("secret-key")

func CreateToken(username string, role string, team string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  username,
		"iss":  "todo-app",
		"aud":  role,
		"Team": team,
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (TokenClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return TokenClaims{}, errors.New("INVALID TOKEN")
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok || !token.Valid {
		return TokenClaims{}, errors.New("INVALID TOKEN")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return TokenClaims{}, errors.New("EXPIRED TOKEN")
	}

	return *claims, nil
}
