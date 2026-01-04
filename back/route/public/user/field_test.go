package users

import (
	"testing"
)

func TestUpdateUserInput_Struct(t *testing.T) {
	input := UpdateUserInput{
		FirstName: "River",
		LastName:  "Smith",
		Password:  "newpassword123",
	}

	if input.FirstName != "River" {
		t.Fatalf("expected FirstName to be River, got %s", input.FirstName)
	}

	if input.LastName != "Smith" {
		t.Fatalf("expected LastName to be Smith, got %s", input.LastName)
	}

	if input.Password != "newpassword123" {
		t.Fatalf("expected Password to be newpassword123, got %s", input.Password)
	}
}

func TestUpdateUserInput_PartialUpdate(t *testing.T) {
	input := UpdateUserInput{
		FirstName: "Doctor",
		LastName:  "",
		Password:  "",
	}

	if input.FirstName != "Doctor" {
		t.Fatalf("expected FirstName to be John, got %s", input.FirstName)
	}

	if input.LastName != "" {
		t.Fatalf("expected LastName to be empty")
	}

	if input.Password != "" {
		t.Fatalf("expected Password to be empty")
	}
}

func TestUserOutput_Struct(t *testing.T) {
	output := UserOutput{
		ID:        "507f1f77bcf86cd799439011",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      "team123",
	}

	if output.ID != "507f1f77bcf86cd799439011" {
		t.Fatalf("expected ID to match")
	}

	if output.FirstName != "Doctor" {
		t.Fatalf("expected FirstName to be John, got %s", output.FirstName)
	}

	if output.LastName != "Who" {
		t.Fatalf("expected LastName to be Who, got %s", output.LastName)
	}

	if output.Email != "Doctor@Tardis.com" {
		t.Fatalf("expected Email to be john@Tardis.com, got %s", output.Email)
	}

	if output.Role != "EMPLOYEE" {
		t.Fatalf("expected Role to be EMPLOYEE, got %s", output.Role)
	}

	if output.Team != "team123" {
		t.Fatalf("expected Team to be team123, got %s", output.Team)
	}
}

func TestUserOutput_NoTeam(t *testing.T) {
	output := UserOutput{
		ID:        "507f1f77bcf86cd799439011",
		FirstName: "River",
		LastName:  "Smith",
		Email:     "jane@Tardis.com",
		Role:      "MANAGER",
		Team:      "",
	}

	if output.Team != "" {
		t.Fatalf("expected Team to be empty string")
	}
}

func TestUserOutput_MultipleRoles(t *testing.T) {
	roles := []string{"EMPLOYEE", "MANAGER", "ADMIN"}
	for i, role := range roles {
		output := UserOutput{
			ID:        "507f1f77bcf86cd79943901" + string(rune(i+1)),
			FirstName: "User",
			LastName:  "Test",
			Email:     "user" + string(rune(i+49)) + "@Tardis.com",
			Role:      role,
			Team:      "",
		}
		if output.Role != role {
			t.Fatalf("expected Role to be %s, got %s", role, output.Role)
		}
	}
}
