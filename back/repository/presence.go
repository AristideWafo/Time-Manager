package repository

import (
	"TimeManager/model"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func SavePresence(presence *model.Presence) error {

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: presence.User}}
	err := GetOneUser(user, filter)

	if err != nil {
		return err
	}

	return Save(presence, PresenceCollection())
}

func UpdateOnePresence(presence *model.Presence, update bson.D) error {

	for _, elem := range update {
		if elem.Key == "User" {
			return errors.New("CANNOT UPDATE PRESENCE")
		}
	}

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: presence.User}}
	err := GetOneUser(user, filter)

	if err != nil {
		return err
	}

	return UpdateOne(presence, UserCollection(), update)
}

func GetManyPresences(presences *[]*model.Presence, filter bson.D) error {

	err := GetMany(filter, PresenceCollection(), presences)

	if err != nil {
		return err
	}

	for _, presence := range *presences {

		user := &model.User{}
		filter := bson.D{{Key: "_id", Value: presence.User}}
		err = GetOneUser(user, filter)

		if err != nil {
			return err
		}

	}

	return nil
}
