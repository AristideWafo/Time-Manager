package repository

import (
	"context"
	"testing"

	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func cleanupUserTest(_ *testing.T) {
	if Client.Database != nil {
		UserCollection().DeleteMany(context.TODO(), bson.D{})
	}
}

func TestGetOneUser_UserWithoutTeam(t *testing.T) {
	setupTestConnection(t)
	defer cleanupUserTest(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      bson.NilObjectID,
	}

	err := SaveUser(user)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	savedID := user.ID

	retrieved := &model.User{}
	filter := bson.D{{Key: "_id", Value: savedID}}
	err = GetOneUser(retrieved, filter)
	if err != nil {
		t.Fatalf("expected no error for GetOneUser without team, got %v", err)
	}

	if retrieved.ID != savedID {
		t.Fatalf("expected retrieved user ID to match saved ID")
	}
}

func TestSaveUser_ValidUser(t *testing.T) {
	setupTestConnection(t)
	defer cleanupUserTest(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := SaveUser(user)
	if err != nil {
		t.Fatalf("expected no error for SaveUser, got %v", err)
	}

	if user.ID.IsZero() {
		t.Fatalf("expected user ID to be set")
	}
}

func TestUpdateOneUser_ValidUpdate(t *testing.T) {
	setupTestConnection(t)
	defer cleanupUserTest(t)

	user := &model.User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
	}

	err := SaveUser(user)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}

	savedID := user.ID
	user.FirstName = "River"

	err = UpdateOneUser(user, bson.D{{Key: "FirstName", Value: "River"}})
	if err != nil {
		t.Fatalf("expected no error for UpdateOneUser, got %v", err)
	}

	if user.ID != savedID {
		t.Fatalf("expected user ID to remain unchanged")
	}

	if user.FirstName != "River" {
		t.Fatalf("expected user FirstName to be updated to River, got %s", user.FirstName)
	}
}
