package model

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestValidateRole_ValidEmployeeRole(t *testing.T) {
	err := ValidateRole("EMPLOYEE")
	if err != nil {
		t.Fatalf("expected no error for EMPLOYEE role, got %v", err)
	}
}

func TestValidateRole_ValidManagerRole(t *testing.T) {
	err := ValidateRole("MANAGER")
	if err != nil {
		t.Fatalf("expected no error for MANAGER role, got %v", err)
	}
}

func TestValidateRole_ValidAdminRole(t *testing.T) {
	err := ValidateRole("ADMIN")
	if err != nil {
		t.Fatalf("expected no error for ADMIN role, got %v", err)
	}
}

func TestValidateRole_InvalidRole(t *testing.T) {
	err := ValidateRole("INVALID")
	if err == nil {
		t.Fatalf("expected error for invalid role, got nil")
	}
}

func TestValidateRole_EmptyRole(t *testing.T) {
	err := ValidateRole("")
	if err == nil {
		t.Fatalf("expected error for empty role, got nil")
	}
}

func TestUser_Validate_ValidUser(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      teamID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		t.Fatalf("expected no error for valid user, got %v", err)
	}
}

func TestUser_Validate_InvalidRole(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "INVALID_ROLE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for invalid role, got nil")
	}
}

func TestUser_Validate_MissingEmail(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for missing email, got nil")
	}
}

func TestUser_Validate_InvalidEmailFormat(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "not-an-email",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for invalid email format, got nil")
	}
}

func TestUser_Validate_MissingPassword(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for missing password, got nil")
	}
}

func TestUser_Validate_MissingFirstName(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for missing FirstName, got nil")
	}
}

func TestUser_Validate_MissingLastName(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err == nil {
		t.Fatalf("expected error for missing LastName, got nil")
	}
}

func TestUser_ValidatePreSave_ValidUser(t *testing.T) {
	now := time.Now()

	user := &User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error for valid user pre-save, got %v", err)
	}
}

func TestUser_ValidatePreSave_InvalidRole(t *testing.T) {
	now := time.Now()

	user := &User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "INVALID_ROLE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.ValidatePreSave()
	if err == nil {
		t.Fatalf("expected error for invalid role in pre-save, got nil")
	}
}

func TestUser_ValidatePreSave_IgnoresID(t *testing.T) {
	now := time.Now()

	user := &User{
		ID:        bson.NilObjectID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.ValidatePreSave()
	if err != nil {
		t.Fatalf("expected no error (ID should be ignored), got %v", err)
	}
}

func TestUser_ValidatePreSave_MissingEmail(t *testing.T) {
	now := time.Now()

	user := &User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.ValidatePreSave()
	if err == nil {
		t.Fatalf("expected error for missing email in pre-save, got nil")
	}
}

func TestUser_ValidatePreSave_InvalidEmailFormat(t *testing.T) {
	now := time.Now()

	user := &User{
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "invalid-email",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.ValidatePreSave()
	if err == nil {
		t.Fatalf("expected error for invalid email format in pre-save, got nil")
	}
}

func TestUser_SetID(t *testing.T) {
	user := &User{}
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")

	user.SetID(id)

	if user.ID != id {
		t.Fatalf("expected ID to be set to %v, got %v", id, user.ID)
	}
}

func TestUser_GetID(t *testing.T) {
	id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	user := &User{ID: id}

	retrievedID := user.GetID()

	if retrievedID != id {
		t.Fatalf("expected ID to be %v, got %v", id, retrievedID)
	}
}

func TestUser_SetCreatedAtAndUpdatedAt(t *testing.T) {
	user := &User{}

	before := time.Now()
	user.SetCreatedAtAndUpdatedAt()
	after := time.Now()

	if user.CreatedAt.IsZero() {
		t.Fatalf("expected CreatedAt to be set, got zero value")
	}
	if user.UpdatedAt.IsZero() {
		t.Fatalf("expected UpdatedAt to be set, got zero value")
	}
	if user.CreatedAt != user.UpdatedAt {
		t.Fatalf("expected CreatedAt and UpdatedAt to be equal, got %v != %v", user.CreatedAt, user.UpdatedAt)
	}

	if user.CreatedAt.Before(before) || user.CreatedAt.After(after) {
		t.Fatalf("expected CreatedAt to be between %v and %v, got %v", before, after, user.CreatedAt)
	}
}

func TestUser_SetCreatedAtAndUpdatedAt_MultipleCalls(t *testing.T) {
	user := &User{}

	user.SetCreatedAtAndUpdatedAt()
	firstTime := user.CreatedAt

	time.Sleep(10 * time.Millisecond)

	user.SetCreatedAtAndUpdatedAt()
	secondTime := user.CreatedAt

	if firstTime.Equal(secondTime) {
		t.Fatalf("expected timestamps to be different after second call")
	}
	if !secondTime.After(firstTime) {
		t.Fatalf("expected second timestamp to be after first")
	}
}

func TestUser_SetID_and_GetID_RoundTrip(t *testing.T) {
	user := &User{}
	originalID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439999")

	user.SetID(originalID)
	retrievedID := user.GetID()

	if retrievedID != originalID {
		t.Fatalf("expected round-trip ID to be %v, got %v", originalID, retrievedID)
	}
}

func TestUser_ValidatePreSave_AllRoles(t *testing.T) {
	roles := []string{"EMPLOYEE", "MANAGER", "ADMIN"}
	now := time.Now()

	for _, role := range roles {
		user := &User{
			Password:  "hashed_password",
			FirstName: "Doctor",
			LastName:  "Who",
			Email:     "Doctor@Tardis.com",
			Role:      role,
			CreatedAt: now,
			UpdatedAt: now,
		}

		err := user.ValidatePreSave()
		if err != nil {
			t.Fatalf("expected no error for role %s, got %v", role, err)
		}
	}
}

func TestUser_Validate_OptionalShiftTimes(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		t.Fatalf("expected no error for user without shift times (optional), got %v", err)
	}
}

func TestUser_Validate_OptionalTeam(t *testing.T) {
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	now := time.Now()

	user := &User{
		ID:        userID,
		Password:  "hashed_password",
		FirstName: "Doctor",
		LastName:  "Who",
		Email:     "Doctor@Tardis.com",
		Role:      "EMPLOYEE",
		Team:      bson.NilObjectID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := user.Validate()
	if err != nil {
		t.Fatalf("expected no error for user without team (optional), got %v", err)
	}
}
