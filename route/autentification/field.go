package autentification

import "github.com/golang-jwt/jwt/v5"

type AutentificationUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
