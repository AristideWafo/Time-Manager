package repository

import "TimeManager/model"

var teams = []model.Team{
	{Name: "Rocket", Users: users},
}

func FetchTeam(name string) (model.Team, bool, error) {
	return teams[0], true, nil
}

func FetchAllTeams() ([]model.Team, bool, error) {
	return teams, true, nil
}

func CreateTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}

func UpdateTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}

func DeleteTeam(team model.Team) (model.Team, bool, error) {
	return team, true, nil
}
