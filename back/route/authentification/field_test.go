package authentification

import (
	"testing"
)

func TestAuthentificationUserInput_Struct(t *testing.T) {
	input := AuthentificationUserInput{
		Email:    "Doctor@Tardis.com",
		Password: "password123",
	}

	if input.Email != "Doctor@Tardis.com" {
		t.Fatalf("expected Email to be john@Tardis.com, got %s", input.Email)
	}

	if input.Password != "password123" {
		t.Fatalf("expected Password to be password123, got %s", input.Password)
	}
}

func TestAuthentificationUserOutput_Struct(t *testing.T) {
	output := AuthentificationUserOutput{
		ID:        "507f1f77bcf86cd799439011",
		Token:     "jwt_token_here",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      "team123",
	}

	if output.ID != "507f1f77bcf86cd799439011" {
		t.Fatalf("expected ID to match")
	}

	if output.Token != "jwt_token_here" {
		t.Fatalf("expected Token to be jwt_token_here, got %s", output.Token)
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

func TestAuthentificationUserOutput_EmptyTeam(t *testing.T) {
	output := AuthentificationUserOutput{
		ID:        "507f1f77bcf86cd799439011",
		Token:     "jwt_token_here",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      "",
	}

	if output.Team != "" {
		t.Fatalf("expected Team to be empty string, got %s", output.Team)
	}
}

func TestAuthentificationUserInput_EmptyFields(t *testing.T) {
	input := AuthentificationUserInput{
		Email:    "",
		Password: "",
	}

	if input.Email != "" {
		t.Fatalf("expected Email to be empty, got %s", input.Email)
	}

	if input.Password != "" {
		t.Fatalf("expected Password to be empty, got %s", input.Password)
	}
}

func TestAuthentificationUserOutput_AllRoles(t *testing.T) {
	roles := []string{"EMPLOYEE", "MANAGER", "ADMIN"}
	for _, role := range roles {
		output := AuthentificationUserOutput{
			ID:        "507f1f77bcf86cd799439011",
			Token:     "jwt_token_here",
			FirstName: "Test",
			LastName:  "User",
			Email:     "test@Tardis.com",
			Role:      role,
			Team:      "team123",
		}
		if output.Role != role {
			t.Fatalf("expected Role to be %s, got %s", role, output.Role)
		}
	}
}
