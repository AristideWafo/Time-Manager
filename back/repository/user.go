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

func UpdateUser(user *model.User, update bson.D) error {

	if user.Team != NULL_ID {
		return Update(user, UserCollection(), update)
	}

	filter := bson.D{{Key: "_id", Value: user.Team}}
	current_team := &model.Team{}
	if err := GetOneTeam(current_team, filter); err != nil {
		return err
	}

	err := Update(user, UserCollection(), update)

	if err != nil {
		return err
	}

	if user.Team != current_team.ID {
		current_team = &model.Team{}
		if err = GetOneTeam(current_team, filter); err != nil {
			return err
		}
	}

	return nil
}

func DeleteUser(user model.User) (model.User, bool, error) {
	return user, true, nil
}
