package repository

import (
	"context"
	"testing"

	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func cleanupTeamTest(_ *testing.T) {
	if Client.Database != nil {
		TeamCollection().DeleteMany(context.TODO(), bson.D{})
	}
}

func TestGetOneTeam_ValidTeam(t *testing.T) {
	setupTestConnection(t)
	defer cleanupTeamTest(t)

	collection := TeamCollection()

	team := &model.Team{
		Name: "Engineering",
	}

	err := Save(team, collection)
	if err != nil {
		t.Fatalf("failed to save team: %v", err)
	}

	savedID := team.ID

	retrieved := &model.Team{}
	filter := bson.D{{Key: "_id", Value: savedID}}
	err = GetOneTeam(retrieved, filter)
	if err != nil {
		t.Fatalf("expected no error for GetOneTeam, got %v", err)
	}

	if retrieved.ID != savedID {
		t.Fatalf("expected retrieved team ID to match saved ID")
	}
}

func TestSaveTeam_ValidTeam(t *testing.T) {
	setupTestConnection(t)
	defer cleanupTeamTest(t)

	team := &model.Team{
		Name: "Engineering",
	}

	err := SaveTeam(team)
	if err != nil {
		t.Fatalf("expected no error for SaveTeam, got %v", err)
	}

	if team.ID.IsZero() {
		t.Fatalf("expected team ID to be set")
	}
}

func TestUpdateOneTeam_ValidUpdate(t *testing.T) {
	setupTestConnection(t)
	defer cleanupTeamTest(t)

	team := &model.Team{
		Name: "Engineering",
	}

	err := SaveTeam(team)
	if err != nil {
		t.Fatalf("failed to save team: %v", err)
	}

	savedID := team.ID
	team.Name = "Product"

	err = UpdateOneTeam(team, bson.D{{Key: "Name", Value: "Product"}})
	if err != nil {
		t.Fatalf("expected no error for UpdateOneTeam, got %v", err)
	}

	if team.ID != savedID {
		t.Fatalf("expected team ID to remain unchanged")
	}

	if team.Name != "Product" {
		t.Fatalf("expected team Name to be updated to Product, got %s", team.Name)
	}
}
