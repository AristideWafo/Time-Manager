package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testPassword    = "ThisIsAPassword"
	badTestPassword = "ThisIsAWrongPassword"
)

func TestHashPassword(t *testing.T) {
	hashed, err := HashPassword(testPassword)
	assert.NoError(t, err)
	assert.Len(t, hashed, 60)
}

func TestCheckPasswordHash(t *testing.T) {
	hash := "$2a$14$aJ7lrRHNeEEjhHkVx6Ldv.p3UpwpADFX8614EJp2cfSV/CPipBf02"

	trueCheck, err := CheckPasswordHash([]byte(testPassword), hash)
	assert.NoError(t, err)
	assert.True(t, trueCheck)

	falseCheck, err := CheckPasswordHash([]byte(badTestPassword), hash)
	assert.NoError(t, err)
	assert.False(t, falseCheck)

}
