package repository

import (
	"TimeManager/model"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOneUser(user *model.User, filter bson.D) error {

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

func GetManyUsers(users *[]*model.User, filter bson.D) error {

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

func SaveUser(user *model.User) error {

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

func UpdateOneUser(user *model.User, update bson.D) error {

	team_update := false
	unset := bson.D{}
	var ok bool

	for index, elem := range update {

		if elem.Key == "Team" && elem.Value != "" {
			team_update = true
			if !ok {
				return errors.New("INVALID TEAM UPDATE")
			}
			break
		}

		if elem.Key == "Team" && elem.Value == "" {
			unset = bson.D{{Key: "Team", Value: ""}}
			update = append(update[:index], update[index+1:]...)
			break
		}
	}

	if !team_update {
		return UpdateOne(user, UserCollection(), update, unset)
	}

	filter := bson.D{{Key: "_id", Value: user.ID}}
	team := &model.Team{}
	if err := GetOneTeam(team, filter); err != nil {
		return err
	}

	return UpdateOne(user, UserCollection(), update, unset)
}

func UpdateManyUsers(filter bson.D, update bson.D) error {

	team_update := false
	unset := bson.D{}
	var team_id bson.ObjectID
	var ok bool

	for index, elem := range update {

		if elem.Key == "Team" && elem.Value != "" {
			team_update = true
			team_id, ok = elem.Value.(bson.ObjectID)
			if !ok {
				return errors.New("INVALID TEAM UPDATE")
			}
			break
		}

		if elem.Key == "Team" && elem.Value == "" {
			unset = bson.D{{Key: "Team", Value: ""}}
			update = append(update[:index], update[index+1:]...)
			break
		}
	}

	user := &model.User{}

	if !team_update {
		return UpdateMany(user, filter, UserCollection(), update, unset)
	}

	team_filter := bson.D{{Key: "_id", Value: team_id}}
	team := &model.Team{}
	if err := GetOneTeam(team, team_filter); err != nil {
		return err
	}

	return UpdateMany(user, filter, UserCollection(), update, unset)
}

func DeleteOneUser(user *model.User) error {
	return DeleteOne(user, UserCollection())
}
