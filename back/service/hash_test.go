package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	hashed, err := HashPassword("ThisIsAPassword")
	assert.NoError(t, err)
	assert.Len(t, hashed, 60)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "ThisIsAPassword"
	badPassword := "ThisIsAWrongPassword"
	hash := "$2a$14$aJ7lrRHNeEEjhHkVx6Ldv.p3UpwpADFX8614EJp2cfSV/CPipBf02"

	trueCheck, err := CheckPasswordHash([]byte(password), hash)
	assert.NoError(t, err)
	assert.True(t, trueCheck)

	falseCheck, err := CheckPasswordHash([]byte(badPassword), hash)
	assert.NoError(t, err)
	assert.False(t, falseCheck)

}
