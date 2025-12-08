package service

import (
	"TimeManager/model"
	"TimeManager/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTeamByName(name string) (*model.Team, error) {
	team := &model.Team{}
	print(name)
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

	return team, repository.UpdateOneTeam(team, update)
}

func GetTeamUsers(_id bson.ObjectID) ([]*model.User, error) {
	list_users := []*model.User{}

	filter := bson.D{{Key: "Team", Value: _id}}
	err := repository.GetManyUsers(&list_users, filter)
	if err != nil {
		return []*model.User{}, err
	}

	if list_users == nil {
		list_users = []*model.User{}
	}

	return list_users, err
}
