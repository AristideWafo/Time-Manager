package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var GetOneUser = func(user *model.User, filter bson.D) error {

	err := GetOne(filter, UserCollection(), user)

	if err != nil {
		return err
	}

	if user.Team == bson.NilObjectID {
		return nil
	}

	filter = bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}
	return GetOneTeam(team, filter)

}

var GetManyUsers = func(users *[]*model.User, filter bson.D) error {

	err := GetMany(filter, UserCollection(), users)

	if err != nil {
		return err
	}

	for _, user := range *users {

		if user.Team == bson.NilObjectID {
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

var SaveUser = func(user *model.User) error {

	if user.Team == bson.NilObjectID {
		return Save(user, UserCollection())
	}

	filter := bson.D{{Key: "_id", Value: user.Team}}
	team := &model.Team{}

	if err := GetOneTeam(team, filter); err != nil {
		return err
	}

	return Save(user, UserCollection())

}

var UpdateOneUser = func(user *model.User, update bson.D) error {

	unset := bson.D{}
	team_id := bson.NilObjectID

	for index, elem := range update {

		if elem.Key == "Team" && elem.Value == bson.NilObjectID {
			unset = bson.D{{Key: "Team", Value: ""}}
			update = append(update[:index], update[index+1:]...)
			break
		}

		if elem.Key == "Team" {
			team_id, _ = elem.Value.(bson.ObjectID)
			break
		}
	}

	if team_id != bson.NilObjectID {
		filter := bson.D{{Key: "_id", Value: team_id}}
		team := &model.Team{}
		if err := GetOneTeam(team, filter); err != nil {
			return err
		}

	}

	return UpdateOne(user, UserCollection(), update, unset)
}

func UpdateManyUsers(filter bson.D, update bson.D) error {

	unset := bson.D{}
	var team_id bson.ObjectID

	for index, elem := range update {

		if elem.Key == "Team" && elem.Value == bson.NilObjectID {
			unset = bson.D{{Key: "Team", Value: ""}}
			update = append(update[:index], update[index+1:]...)
			break
		}

		if elem.Key == "Team" {
			team_id, _ = elem.Value.(bson.ObjectID)
			break
		}
	}

	if team_id != bson.NilObjectID {
		filter := bson.D{{Key: "_id", Value: team_id}}
		team := &model.Team{}
		if err := GetOneTeam(team, filter); err != nil {
			return err
		}
	}

	users := []*model.User{}

	if err := GetMany(filter, UserCollection(), &users); err != nil {
		return err
	}

	return UpdateMany(&users, filter, UserCollection(), update, unset)
}

func DeleteOneUser(user *model.User) error {
	return DeleteOne(user, UserCollection())
}
