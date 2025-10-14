package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashed, err := bcrypt.GenerateFromPassword(passwordBytes, 14)

	for i := range passwordBytes {
		passwordBytes[i] = 0
	}
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func CheckPasswordHash(password []byte, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	for i := range password {
		password[i] = 0
	}
	if err != nil {
		return false, err
	}
	return err == nil, nil
}
