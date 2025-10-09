package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// TO DO: DUMMY REPLIES

var NULL_ID, _ = bson.ObjectIDFromHex("")

func Login(email string, password string) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "Email", Value: email}, {Key: "Password", Value: password}}
	err := GetOne(filter, UserCollection(), user)

	return user, err
}

func GetUser(_id bson.ObjectID) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: _id}}
	err := GetOne(filter, UserCollection(), user)

	if err != nil {
		return user, err
	}

	if user.Team == NULL_ID {
		return user, nil
	}

	filter = bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}
	err = GetOne(filter, TeamCollection(), team)

	return user, err
}

func CreateUser(first_name string, last_name string, email string, password string, role string) (*model.User, error) {
	user := &model.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
		Role:      role,
	}

	err := Save(user, UserCollection())

	return user, err
}

func UpdateUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}

func DeleteUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}
