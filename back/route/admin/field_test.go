package admin

import (
	"testing"
)

func TestCreateUserInput_Struct(t *testing.T) {
	input := CreateUserInput{
		FirstName: "Doctor",
		LastName:  "Who",
		Role:      "EMPLOYEE",
		Email:     "Doctor@Tardis.com",
		Password:  "password123",
		Team:      "team123",
	}

	if input.FirstName != "Doctor" {
		t.Fatalf("expected FirstName to be John, got %s", input.FirstName)
	}

	if input.LastName != "Who" {
		t.Fatalf("expected LastName to be Who, got %s", input.LastName)
	}

	if input.Role != "EMPLOYEE" {
		t.Fatalf("expected Role to be EMPLOYEE, got %s", input.Role)
	}

	if input.Email != "Doctor@Tardis.com" {
		t.Fatalf("expected Email to be john@Tardis.com, got %s", input.Email)
	}

	if input.Password != "password123" {
		t.Fatalf("expected Password to be password123, got %s", input.Password)
	}

	if input.Team != "team123" {
		t.Fatalf("expected Team to be team123, got %s", input.Team)
	}
}

func TestUpdateUserTeamInput_Struct(t *testing.T) {
	input := UpdateUserTeamInput{
		TeamId: "team456",
	}

	if input.TeamId != "team456" {
		t.Fatalf("expected TeamId to be team456, got %s", input.TeamId)
	}
}

func TestCreateTeamInput_Struct(t *testing.T) {
	input := CreateTeamInput{
		Name: "Engineering",
	}

	if input.Name != "Engineering" {
		t.Fatalf("expected Name to be Engineering, got %s", input.Name)
	}
}

func TestUpdateTeamInput_Struct(t *testing.T) {
	input := UpdateTeamInput{
		CurrentName: "OldTeam",
		NewName:     "NewTeam",
	}

	if input.CurrentName != "OldTeam" {
		t.Fatalf("expected CurrentName to be OldTeam, got %s", input.CurrentName)
	}

	if input.NewName != "NewTeam" {
		t.Fatalf("expected NewName to be NewTeam, got %s", input.NewName)
	}
}

func TestRequestedTeamInput_Struct(t *testing.T) {
	input := RequestedTeamInput{
		Name: "TeamName",
	}

	if input.Name != "TeamName" {
		t.Fatalf("expected Name to be TeamName, got %s", input.Name)
	}
}

func TestCreateUserInput_AllRoles(t *testing.T) {
	roles := []string{"EMPLOYEE", "MANAGER", "ADMIN"}
	for _, role := range roles {
		input := CreateUserInput{
			FirstName: "Test",
			LastName:  "User",
			Role:      role,
			Email:     "test@Tardis.com",
			Password:  "password",
			Team:      "",
		}
		if input.Role != role {
			t.Fatalf("expected Role to be %s, got %s", role, input.Role)
		}
	}
}
