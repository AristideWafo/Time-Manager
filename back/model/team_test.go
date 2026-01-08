package model

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestTeam_Validate_ValidTeam(t *testing.T) {
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	team := &Team{
		ID:        teamID,
		Name:      "Engineering",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := team.Validate()
	if err != nil {
		t.Fatalf("expected no error for valid team, got %v", err)
	}
}

func TestTeam_Validate_MissingName(t *testing.T) {
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	team := &Team{
		ID:        teamID,
		Name:      "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := team.Validate()
	if err == nil {
		t.Fatalf("expected error for missing Name, got nil")
	}
}

func TestTeam_Validate_MissingCreatedAt(t *testing.T) {
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")

	team := &Team{
		ID:        teamID,
		Name:      "Engineering",
		CreatedAt: time.Time{},
		UpdatedAt: time.Now(),
	}

	err := team.Validate()
	if err == nil {
		t.Fatalf("expected error for missing CreatedAt, got nil")
	}
}

func TestTeam_Validate_MissingUpdatedAt(t *testing.T) {
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")

	team := &Team{
		ID:        teamID,
		Name:      "Engineering",
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	err := team.Validate()
	if err == nil {
		t.Fatalf("expected error for missing UpdatedAt, got nil")
	}
}

func TestTeam_ValidatePreSave_ValidTeam(t *testing.T) {
	now := time.Now()

	team := &Team{
		Name:      "Engineering",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := team.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error for valid team pre-save, got %v", err)
	}
}

func TestTeam_ValidatePreSave_MissingName(t *testing.T) {
	now := time.Now()

	team := &Team{
		Name:      "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := team.ValidatePreSave()
	if err == nil {
		t.Fatalf("expected error for missing Name in pre-save, got nil")
	}
}

func TestTeam_ValidatePreSave_IgnoresID(t *testing.T) {
	now := time.Now()

	team := &Team{
		ID:        bson.NilObjectID,
		Name:      "Engineering",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := team.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error (ID should be ignored), got %v", err)
	}
}

func TestTeam_SetID(t *testing.T) {
	team := &Team{}
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")

	team.SetID(id)

	if team.ID != id {
		t.Fatalf("expected ID to be set to %v, got %v", id, team.ID)
	}
}

func TestTeam_GetID(t *testing.T) {
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	team := &Team{ID: id}

	retrievedID := team.GetID()

	if retrievedID != id {
		t.Fatalf("expected ID to be %v, got %v", id, retrievedID)
	}
}

func TestTeam_SetCreatedAtAndUpdatedAt(t *testing.T) {
	team := &Team{}

	before := time.Now()
	team.SetCreatedAtAndUpdatedAt()
	after := time.Now()

	if team.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set, got zero value")
	}
	if team.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set, got zero value")
	}
	if team.CreatedAt != team.UpdatedAt {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal, got %v != %v", team.CreatedAt, team.UpdatedAt)
	}

	if team.CreatedAt.Before(before) || team.CreatedAt.After(after) {
		t.Fatalf("expected CreatedAt to be between %v and %v, got %v", before, after, team.CreatedAt)
	}
}

func TestTeam_SetCreatedAtAndUpdatedAt_MultipleCalls(t *testing.T) {
	team := &Team{}

	team.SetCreatedAtAndUpdatedAt()
	firstTime := team.CreatedAt

	time.Sleep(10 * time.Millisecond)

	team.SetCreatedAtAndUpdatedAt()
	secondTime := team.CreatedAt

	if firstTime.Equal(secondTime) {
		t.Fatalf("expected timestamps to be different after second call")
	}
	if !secondTime.After(firstTime) {
		t.Fatalf("expected second timestamp to be after first")
	}
}

func TestTeam_SetID_and_GetID_RoundTrip(t *testing.T) {
	team := &Team{}
	originalID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")

	team.SetID(originalID)
	retrievedID := team.GetID()

	if retrievedID != originalID {
		t.Fatalf("expected round-trip ID to be %v, got %v", originalID, retrievedID)
	}
}
