package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TO DO: DUMMY REPLIES

var _id, _ = primitive.ObjectIDFromHex("68e8c75630834f60973523e8")

var users = []model.User{
	{ID: _id, Email: "Blue@IHateEpitech.com", Password: "Red", Role: "ADMIN", Team: "Rocket", FirstName: "Miel", LastName: "Pops"},
}

func Login(email string, password string) (model.User, bool, error) {
	if email != users[0].Email || password != users[0].Password {
		return model.User{}, false, nil
	}
	return users[0], true, nil
}

func GetUser(_id primitive.ObjectID) (model.User, bool, error) {
	if _id != users[0].ID {
		return model.User{}, false, nil
	}
	return users[0], true, nil
}

func GetAllUsers() ([]model.User, bool, error) {
	return users, true, nil
}

func CreateUser(email string, password string) (model.User, bool, error) {
	return users[0], true, nil
}

func UpdateUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}

func DeleteUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}
