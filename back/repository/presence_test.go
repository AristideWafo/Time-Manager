package repository

import (
	"context"
	"testing"
	"time"

	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func cleanupPresenceTest(t *testing.T) {
	if Client.Database != nil {
		_, err := UserCollection().DeleteMany(context.TODO(), bson.D{})

		if err != nil {
			t.Fatalf("failed to cleanup test collection: %v", err)
		}
		_, err = PresenceCollection().DeleteMany(context.TODO(), bson.D{})

		if err != nil {
			t.Fatalf("failed to cleanup test collection: %v", err)
		}
	}
}

func TestSavePresence_ValidPresence(t *testing.T) {
	setupTestConnection(t)
	defer cleanupPresenceTest(t)

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

	presence := &model.Presence{
		Type:      "PRESENT",
		Timestamp: time.Now(),
		User:      user.ID,
	}

	err = SavePresence(presence)
	if err != nil {
		t.Fatalf("expected no error for SavePresence, got %v", err)
	}

	if presence.ID.IsZero() {
		t.Fatalf("expected presence ID to be set")
	}
}

func TestSavePresence_InvalidUser(t *testing.T) {
	setupTestConnection(t)
	defer cleanupPresenceTest(t)

	nonExistentUserID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")
	presence := &model.Presence{
		Type:      "PRESENT",
		Timestamp: time.Now(),
		User:      nonExistentUserID,
	}

	err := SavePresence(presence)
	if err == nil {
		t.Fatalf("expected error for SavePresence with invalid user, got nil")
	}
}

func TestGetManyPresences_ValidUser(t *testing.T) {
	setupTestConnection(t)
	defer cleanupPresenceTest(t)

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

	presence1 := &model.Presence{
		Type:      "PRESENT",
		Timestamp: time.Now(),
		User:      user.ID,
	}

	presence2 := &model.Presence{
		Type:      "ABSENT",
		Timestamp: time.Now(),
		User:      user.ID,
	}

	err = SavePresence(presence1)
	if err != nil {
		t.Fatalf("failed to save presence1: %v", err)
	}

	err = SavePresence(presence2)
	if err != nil {
		t.Fatalf("failed to save presence2: %v", err)
	}

	retrieved := make([]*model.Presence, 0)
	filter := bson.D{{Key: "User", Value: user.ID}}
	err = GetManyPresences(&retrieved, filter)
	if err != nil {
		t.Fatalf("expected no error for GetManyPresences, got %v", err)
	}

	if len(retrieved) < 2 {
		t.Fatalf("expected at least 2 presences, got %d", len(retrieved))
	}
}
