package service

import (
	"TimeManager/model"
	"TimeManager/repository"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

var team_name = "MUDA MUDA"

func TestGetTeamByName_Success(t *testing.T) {
	orig := repository.GetOneTeam
	defer func() { repository.GetOneTeam = orig }()

	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		team.Name = team_name
		return nil
	}

	team, err := GetTeamByName(team_name)

	if team.Name != team_name {
		t.Fatalf("expected team %s, got %s", team_name, team.Name)
	}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}

func TestGetTeamByName_NotFound(t *testing.T) {
	orig := repository.GetOneTeam
	defer func() { repository.GetOneTeam = orig }()

	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		return errors.New("not found")
	}

	_, err := GetTeamByName("does-not-exist")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetTeamById_Success(t *testing.T) {
	orig := repository.GetOneTeam
	defer func() { repository.GetOneTeam = orig }()

	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439099")

	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		team.ID = teamID
		team.Name = team_name
		return nil
	}

	team, err := GetTeamById(teamID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if team.ID != teamID || team.Name != team_name {
		t.Fatalf("unexpected team: %+v", team)
	}
}

func TestCreateTeam_Success(t *testing.T) {
	orig := repository.SaveTeam
	defer func() { repository.SaveTeam = orig }()

	var saved *model.Team
	repository.SaveTeam = func(team *model.Team) error {
		saved = team
		// emulate DB setting an ID
		id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd7994390aa")
		team.ID = id
		return nil
	}

	team, err := CreateTeam(team_name)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if saved == nil {
		t.Fatalf("expected SaveTeam to be called")
	}
	if team.Name != team_name {
		t.Fatalf("expected team name %s, got %s", team_name, team.Name)
	}
}

func TestCreateTeam_SaveError(t *testing.T) {
	orig := repository.SaveTeam
	defer func() { repository.SaveTeam = orig }()

	repository.SaveTeam = func(team *model.Team) error {
		return errors.New("db fail")
	}

	_, err := CreateTeam(team_name)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestUpdateTeam_Success(t *testing.T) {
	origGet := repository.GetOneTeam
	origUpdate := repository.UpdateOneTeam
	defer func() { repository.GetOneTeam = origGet; repository.UpdateOneTeam = origUpdate }()

	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		team.Name = team_name
		return nil
	}

	repository.UpdateOneTeam = func(team *model.Team, update bson.D) error {
		// emulate applying update: set Name if present
		for _, e := range update {
			if e.Key == "Name" {
				if v, ok := e.Value.(string); ok {
					team.Name = v
				}
			}
		}
		return nil
	}

	newName := "NEW TEAM"
	updated, err := UpdateTeam(team_name, bson.D{{Key: "Name", Value: newName}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Name != newName {
		t.Fatalf("expected updated name %s, got %s", newName, updated.Name)
	}
}

func TestUpdateTeam_NotFound(t *testing.T) {
	orig := repository.GetOneTeam
	defer func() { repository.GetOneTeam = orig }()

	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		return errors.New("not found")
	}

	_, err := UpdateTeam("nope", bson.D{{Key: "Name", Value: "x"}})
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestGetTeamUsers_Success(t *testing.T) {
	orig := repository.GetManyUsers
	defer func() { repository.GetManyUsers = orig }()

	repository.GetManyUsers = func(users *[]*model.User, filter bson.D) error {
		*users = []*model.User{{FirstName: "A"}}
		return nil
	}

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd7994390bb")
	list, err := GetTeamUsers(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 1 || list[0].FirstName != "A" {
		t.Fatalf("unexpected list: %+v", list)
	}
}

func TestGetTeamUsers_Error(t *testing.T) {
	orig := repository.GetManyUsers
	defer func() { repository.GetManyUsers = orig }()

	repository.GetManyUsers = func(users *[]*model.User, filter bson.D) error {
		return errors.New("db fail")
	}

	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd7994390bb")
	_, err := GetTeamUsers(id)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
