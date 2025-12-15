package service

import "go.mongodb.org/mongo-driver/v2/bson"

func IsUserInTeam(_id bson.ObjectID, team_id bson.ObjectID) (bool, error) {

	user, err := GetUserByID(_id)

	if err != nil {
		return false, err
	}

	if user.Team != team_id {
		return false, nil
	}

	return true, nil
}
