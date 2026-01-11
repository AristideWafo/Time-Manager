package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var GetOneTeam = func(team *model.Team, filter bson.D) error {
	return GetOne(filter, TeamCollection(), team)
}

var GetManyTeams = func(teams *[]*model.Team, filter bson.D) error {
	return GetMany(filter, TeamCollection(), teams)
}

var SaveTeam = func(team *model.Team) error {
	return Save(team, TeamCollection())
}

var UpdateOneTeam = func(team *model.Team, update bson.D) error {
	return UpdateOne(team, TeamCollection(), update, bson.D{})
}

func DeleteOneTeam(team *model.Team) error {

	err := DeleteOne(team, TeamCollection())

	if err != nil {
		return err
	}

	return nil
}
