package repository

import (
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	t.Run("Should return user if not nil", func(t *testing.T) {
	})
}

// func Login(email string, password string) (*model.User, error) {

// 	user := &model.User{}
// 	filter := bson.D{{Key: "Email", Value: email}, {Key: "Password", Value: password}}
// 	err := GetOne(filter, UserCollection(), user)

// 	return user, err
// }
