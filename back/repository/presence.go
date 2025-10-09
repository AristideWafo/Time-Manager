package repository

import (
	"TimeManager/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreatePresence(in_or_out string, user_id bson.ObjectID, timestamp time.Time) (*model.Presence, error) {

	user := &model.User{}
	filter := bson.D{{Key: "_id", Value: user_id}}
	err := GetOne(filter, UserCollection(), user)

	if err != nil {
		return &model.Presence{}, err
	}

	presence := &model.Presence{
		Type:      in_or_out,
		Timestamp: timestamp,
		User:      user_id,
	}

	err = Save(Document(presence), PresenceCollection())

	return presence, err
}

func GetAllPresencesByUser(user_id bson.ObjectID) (*[]*model.Presence, error) {

	presences := &[]*model.Presence{}
	filter := bson.D{{Key: "User", Value: user_id}}

	err := GetMany(filter, PresenceCollection(), presences)

	return presences, err
}
