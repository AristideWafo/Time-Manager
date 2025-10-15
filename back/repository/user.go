package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOneUser(user *model.User, filter bson.D) error {

	err := GetOne(filter, UserCollection(), user)

	if err != nil {
		return err
	}

	if user.Team == NULL_ID {
		return nil
	}

	filter = bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}
	return GetOneTeam(team, filter)

}

func GetManyUsers(users *[]*model.User, filter bson.D) error {

	err := GetMany(filter, UserCollection(), users)

	if err != nil {
		return err
	}

	for _, user := range *users {

		if user.Team == NULL_ID {
			continue
		}
		filter = bson.D{{Key: "_id", Value: user.Team}}
		team := &model.Team{}
		err = GetOneTeam(team, filter)

		if err != nil {
			return err
		}
	}

	return nil

}

func SaveUser(user *model.User) error {

	if user.Team == NULL_ID {
		return Save(user, UserCollection())
	}

	filter := bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}

	if err := GetOneTeam(team, filter); err != nil {
		return err
	}

	return Save(user, UserCollection())

}

func UpdateOneUser(user *model.User, update bson.D) error {

	ok := false

	for _, elem := range update {

		if elem.Key == "Team" && elem.Value != nil {
			ok = true
			break
		}
	}

	if !ok {
		return UpdateOne(user, UserCollection(), update)
	}

	filter := bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}
	if err := GetOneTeam(team, filter); err != nil {
		return err
	}

	return UpdateOne(user, UserCollection(), update)
}

func UpdateManyUsers(filter bson.D, update bson.D) error {
	ok := false

	for _, elem := range update {

		if elem.Key == "Team" && elem.Value != nil {
			ok = true
			break
		}
	}

	user := &model.User{}

	if !ok {
		return UpdateMany(user, filter, UserCollection(), update)
	}

	team_filter := bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}
	if err := GetOneTeam(team, team_filter); err != nil {
		return err
	}

	return UpdateMany(user, filter, UserCollection(), update)
}

func DeleteOneUser(user *model.User) error {
	return DeleteOne(user, UserCollection())
}
