package repository

import "testApi/model"

var users = []model.User{
	{Username: "Blue", Password: "Red", Role: "ADMIN", Team: "Rocket"},
}

func FetchUser(username string, password string) (model.User, bool, error) {
	return users[0], true, nil
}
