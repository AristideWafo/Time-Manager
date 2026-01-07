package repository

import (
	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOneTeam(team *model.Team, filter bson.D) error {
	return GetOne(filter, TeamCollection(), team)
}

func GetManyTeams(teams *[]*model.Team, filter bson.D) error {
	return GetMany(filter, TeamCollection(), teams)
}

func SaveTeam(team *model.Team) error {
	return Save(team, TeamCollection())
}

func UpdateOneTeam(team *model.Team, update bson.D) error {
	return UpdateOne(team, TeamCollection(), update, bson.D{})
}

func DeleteOneTeam(team *model.Team) error {

	err := DeleteOne(team, TeamCollection())

	if err != nil {
		return err
	}

	return nil
}
