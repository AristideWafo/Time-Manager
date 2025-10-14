package service

import (
	"TimeManager/model"
	"TimeManager/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreatePresence(in_or_out string, user_id bson.ObjectID, timestamp time.Time) (*model.Presence, error) {

	presence := &model.Presence{
		Type:      in_or_out,
		Timestamp: timestamp,
		User:      user_id,
	}

	return presence, repository.SavePresence(presence)
}

func GetAllPresencesByUserID(user_id bson.ObjectID) (*[]*model.Presence, error) {

	presences := &[]*model.Presence{}
	filter := bson.D{{Key: "User", Value: user_id}}

	return presences, repository.GetManyPresences(presences, filter)
}
