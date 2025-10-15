package service

import (
	"TimeManager/model"
	"TimeManager/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Login(email string, password string) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "Email", Value: email}, {Key: "Password", Value: password}}
	return user, repository.GetOneUser(user, filter)
}

func GetUserByID(_id bson.ObjectID) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: _id}}

	return user, repository.GetOneUser(user, filter)
}

func CreateUser(first_name string, last_name string, email string, password string, role string, team_name string) (*model.User, error) {

	user := &model.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
		Role:      role,
	}

	if team_name == "" {
		return user, repository.SaveUser(user)
	}

	team := &model.Team{}
	filter := bson.D{{Key: "Name", Value: team_name}}

	err := repository.GetOneTeam(team, filter)

	if err != nil {
		return &model.User{}, err
	}

	user.Team = team.ID

	return user, repository.SaveUser(user)
}

func UpdateUserByID(_id bson.ObjectID, update bson.D) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: _id}}
	err := repository.GetOneUser(user, filter)

	if err != nil {
		return user, err
	}

	return user, repository.UpdateOneUser(user, update)
}
