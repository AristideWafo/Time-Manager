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

func UpdateTeam(team *model.Team, update bson.D) error {
	return Update(team, TeamCollection(), update)
}

func DeleteTeam(team model.Team) error {
	return nil
}
