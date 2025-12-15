package service

import (
	"TimeManager/model"
	"TimeManager/repository"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Login(email string, password string) (*model.User, error) {

	user := &model.User{}
	filter := bson.D{{Key: "Email", Value: email}}

	err := repository.GetOneUser(user, filter)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	check, err := CheckPasswordHash([]byte(password), user.Password)

	if !check {
		return &model.User{}, errors.New("incorrect password")
	}

	if err != nil {
		return &model.User{}, err
	}

	return user, err
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

	password, err := HashPassword(password)

	if err != nil {
		return &model.User{}, err
	}

	if password == "" {
		return &model.User{}, errors.New("hashing failed, password returned empty")
	}

	user.Password = password

	if team_name == "" {
		return user, repository.SaveUser(user)
	}

	team := &model.Team{}
	filter := bson.D{{Key: "Name", Value: team_name}}

	err = repository.GetOneTeam(team, filter)

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

	for i := range update {
		if update[i].Key == "Password" {
			if password, ok := update[i].Value.(string); ok {
				update[i].Value, err = HashPassword(password)
				if err != nil {
					return user, err
				}
			} else if !ok {
				return user, errors.New("error : password is not a string")
			}
		}

		if update[i].Key == "Team" {
			if team_id, ok := update[i].Value.(string); ok {
				update[i].Value, err = bson.ObjectIDFromHex(team_id)
				if err != nil {
					return user, err
				}
			} else if !ok {
				return user, errors.New("error : invalid team Id")
			}
		}
	}

	if err != nil {
		return user, err
	}

	return user, repository.UpdateOneUser(user, update)
}
