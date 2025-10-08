package repository

import "testApi/model"

var users = []model.User{
	{Username: "Blue", Password: "Red", Role: "ADMIN", Team: "Rocket"},
}

func Login(username string, password string) (model.User, bool, error) {
	return users[0], true, nil
}

func FetchUser(username string) (model.User, bool, error) {
	return users[0], true, nil
}

func FetchAllUsers() ([]model.User, bool, error) {
	return users, true, nil
}

func CreateUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}

func UpdateUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}

func DeleteUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}
