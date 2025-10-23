package service

import (
	"TimeManager/model"
	"TimeManager/repository"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTeamByName(name string) (*model.Team, error) {
	team := &model.Team{}
	filter := bson.D{{Key: "Name", Value: name}}

	return team, repository.GetOneTeam(team, filter)
}

func CreateTeam(name string) (*model.Team, error) {

	team := &model.Team{
		Name: name,
	}
	return team, repository.SaveTeam(team)
}

func UpdateTeam(name string, update bson.D) (*model.Team, error) {

	team := &model.Team{}

	filter := bson.D{{Key: "Name", Value: name}}
	err := repository.GetOneTeam(team, filter)

	if err != nil {
		return team, err
	}

	fmt.Printf("%v", update)

	return team, repository.UpdateOneTeam(team, update)
}
