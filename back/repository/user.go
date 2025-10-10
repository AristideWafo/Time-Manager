package repository

import (
	"TimeManager/model"
)

// TO DO: DUMMY REPLIES

var users = []model.User{
	{Username: "Blue", Password: "Red", Role: "ADMIN", Team: "Rocket"},
}

func Login(username string, password string) (model.User, bool, error) {
	return users[0], true, nil
}

func GetUser(username string) (model.User, bool, error) {
	return users[0], true, nil
}

func GetAllUsers() ([]model.User, bool, error) {
	return users, true, nil
}

func CreateUser(username string, password string) (model.User, bool, error) {
	return users[0], true, nil
}

func UpdateUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}

func DeleteUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}
