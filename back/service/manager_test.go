package service

import (
	"errors"
	"testing"

	"TimeManager/model"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestIsUserInTeam_UserInTeam(t *testing.T) {
	orig := GetUserByID
	defer func() { GetUserByID = orig }()

	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")

	GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		if _id != userID {
			return &model.User{}, nil
		}
		return &model.User{ID: userID, Team: teamID}, nil
	}

	inTeam, err := IsUserInTeam(userID, teamID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !inTeam {
		t.Fatalf("expected user to be in team")
	}
}

func TestIsUserInTeam_UserNotInTeam(t *testing.T) {
	orig := GetUserByID
	defer func() { GetUserByID = orig }()

	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	otherTeamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439013")

	GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{ID: userID, Team: otherTeamID}, nil
	}

	inTeam, err := IsUserInTeam(userID, teamID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if inTeam {
		t.Fatalf("expected user NOT to be in team")
	}
}

func TestIsUserInTeam_GetUserError(t *testing.T) {
	orig := GetUserByID
	defer func() { GetUserByID = orig }()

	targetTeamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")

	expectedErr := errors.New("db error")
	GetUserByID = func(_id bson.ObjectID) (*model.User, error) {
		return &model.User{}, expectedErr
	}

	inTeam, err := IsUserInTeam(userID, targetTeamID)
	if err == nil {
		t.Fatalf("expected an error but got nil")
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("unexpected error: %v", err)
	}
	if inTeam {
		t.Fatalf("expected inTeam to be false when error occurs")
	}
}
