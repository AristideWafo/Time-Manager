package model

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPresence_Validate_ValidPresence(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	presenceID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	now := time.Now()

	presence := &Presence{
		ID:        presenceID,
		Type:      "PRESENT",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.Validate()
	if err != nil {
		t.Fatalf("expected no error for valid presence, got %v", err)
	}
}

func TestPresence_Validate_MissingType(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	presenceID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	now := time.Now()

	presence := &Presence{
		ID:        presenceID,
		Type:      "",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.Validate()
	if err == nil {
		t.Fatalf("expected error for missing Type, got nil")
	}
}

func TestPresence_Validate_MissingUser(t *testing.T) {
	presenceID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	now := time.Now()

	presence := &Presence{
		ID:        presenceID,
		Type:      "PRESENT",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      bson.NilObjectID,
	}

	err := presence.Validate()
	if err == nil {
		t.Fatalf("expected error for missing User, got nil")
	}
}

func TestPresence_Validate_MissingTimestamp(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	presenceID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	now := time.Now()

	presence := &Presence{
		ID:        presenceID,
		Type:      "PRESENT",
		Timestamp: time.Time{},
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.Validate()
	if err == nil {
		t.Fatalf("expected error for missing Timestamp, got nil")
	}
}

func TestPresence_ValidatePreSave_ValidPresence(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	presence := &Presence{
		Type:      "PRESENT",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error for valid presence pre-save, got %v", err)
	}
}

func TestPresence_ValidatePreSave_MissingType(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	presence := &Presence{
		Type:      "",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.ValidatePreSave()
	if err == nil {
		t.Fatalf("expected error for missing Type in pre-save, got nil")
	}
}

func TestPresence_ValidatePreSave_IgnoresID(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	presence := &Presence{
		ID:        bson.NilObjectID,
		Type:      "PRESENT",
		Timestamp: now,
		CreatedAt: now,
		UpdatedAt: now,
		User:      userID,
	}

	err := presence.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error (ID should be ignored), got %v", err)
	}
}

func TestPresence_SetID(t *testing.T) {
	presence := &Presence{}
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")

	presence.SetID(id)

	if presence.ID != id {
		t.Fatalf("expected ID to be set to %v, got %v", id, presence.ID)
	}
}

func TestPresence_GetID(t *testing.T) {
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	presence := &Presence{ID: id}

	retrievedID := presence.GetID()

	if retrievedID != id {
		t.Fatalf("expected ID to be %v, got %v", id, retrievedID)
	}
}

func TestPresence_SetCreatedAtAndUpdatedAt(t *testing.T) {
	presence := &Presence{}

	before := time.Now()
	presence.SetCreatedAtAndUpdatedAt()
	after := time.Now()

	if presence.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set, got zero value")
	}
	if presence.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set, got zero value")
	}
	if presence.CreatedAt != presence.UpdatedAt {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal, got %v != %v", presence.CreatedAt, presence.UpdatedAt)
	}

	if presence.CreatedAt.Before(before) || presence.CreatedAt.After(after) {
		t.Fatalf("expected CreatedAt to be between %v and %v, got %v", before, after, presence.CreatedAt)
	}
}

func TestPresence_SetCreatedAtAndUpdatedAt_MultipleCalls(t *testing.T) {
	presence := &Presence{}

	presence.SetCreatedAtAndUpdatedAt()
	firstTime := presence.CreatedAt

	time.Sleep(10 * time.Millisecond)

	presence.SetCreatedAtAndUpdatedAt()
	secondTime := presence.CreatedAt

	if firstTime.Equal(secondTime) {
		t.Fatalf("expected timestamps to be different after second call")
	}
	if !secondTime.After(firstTime) {
		t.Fatalf("expected second timestamp to be after first")
	}
}

func TestPresence_SetID_and_GetID_RoundTrip(t *testing.T) {
	presence := &Presence{}
	originalID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")

	presence.SetID(originalID)
	retrievedID := presence.GetID()

	if retrievedID != originalID {
		t.Fatalf("expected round-trip ID to be %v, got %v", originalID, retrievedID)
	}
}
