package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TokenClaims struct {
	Team string `json:"team" binding:"required" validate:"required"`
	jwt.RegisteredClaims
}

func CreateToken(_id bson.ObjectID, role string, team bson.ObjectID) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  _id.Hex(),
		"iss":  "todo-app",
		"aud":  role,
		"team": team.Hex(),
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := claims.SignedString(secretKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (TokenClaims, error) {

	secretKey := []byte(os.Getenv("SECRET_KEY"))
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
