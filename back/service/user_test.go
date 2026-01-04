package service

import (
	"TimeManager/model"
	"TimeManager/repository"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestLogin_Success(t *testing.T) {
	orig := repository.GetOneUser
	defer func() { repository.GetOneUser = orig }()

	password := "secret"
	hash, _ := HashPassword(password)

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		user.Email = "a@Tardis.com"
		user.Password = hash
		return nil
	}

	u, err := Login("a@Tardis.com", password)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u == nil || u.Email != "a@Tardis.com" {
		t.Fatalf("unexpected user: %+v", u)
	}
}

func TestLogin_IncorrectPassword(t *testing.T) {
	orig := repository.GetOneUser
	defer func() { repository.GetOneUser = orig }()

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		user.Email = "a@Tardis.com"
		user.Password = "not-a-hash"
		return nil
	}

	u, err := Login("a@Tardis.com", "wrong")
	if err == nil {
		t.Fatalf("expected error for incorrect password")
	}
	if u == nil {
		t.Fatalf("expected user pointer, got nil")
	}
}

func TestLogin_GetUserError(t *testing.T) {
	orig := repository.GetOneUser
	defer func() { repository.GetOneUser = orig }()

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		return errors.New("db fail")
	}

	_, err := Login("a@Tardis.com", "x")
	if err == nil {
		t.Fatalf("expected error when GetOneUser fails")
	}
}

func TestCreateUser_NoTeam_Success(t *testing.T) {
	origSave := repository.SaveUser
	defer func() { repository.SaveUser = origSave }()

	var saved *model.User
	repository.SaveUser = func(user *model.User) error {
		saved = user
		id, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439010")
		user.ID = id
		return nil
	}

	u, err := CreateUser("First", "Last", "a@b.c", "pwd", "EMPLOYEE", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if saved == nil {
		t.Fatalf("expected SaveUser to be called")
	}
	if u.Email != "a@b.c" {
		t.Fatalf("unexpected user: %+v", u)
	}
}

func TestCreateUser_SaveError(t *testing.T) {
	origSave := repository.SaveUser
	defer func() { repository.SaveUser = origSave }()

	repository.SaveUser = func(user *model.User) error {
		return errors.New("save fail")
	}

	_, err := CreateUser("F", "L", "a@b.c", "pwd", "EMPLOYEE", "")
	if err == nil {
		t.Fatalf("expected error when save fails")
	}
}

func TestCreateUser_WithTeam(t *testing.T) {
	origGetTeam := repository.GetOneTeam
	origSave := repository.SaveUser
	defer func() { repository.GetOneTeam = origGetTeam; repository.SaveUser = origSave }()

	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439020")
	repository.GetOneTeam = func(team *model.Team, filter bson.D) error {
		team.ID = teamID
		team.Name = "T"
		return nil
	}

	repository.SaveUser = func(user *model.User) error {
		return nil
	}

	u, err := CreateUser("F", "L", "a@b.c", "pwd", "EMPLOYEE", "T")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Team != teamID {
		t.Fatalf("expected team %v, got %v", teamID, u.Team)
	}
}

func TestUpdateUserByID_PasswordAndTeam(t *testing.T) {
	origGet := repository.GetOneUser
	origUpdate := repository.UpdateOneUser
	defer func() { repository.GetOneUser = origGet; repository.UpdateOneUser = origUpdate }()

	uid, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439030")
	teamID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439031")

	repository.GetOneUser = func(user *model.User, filter bson.D) error {
		user.ID = uid
		user.Email = "a@b.c"
		return nil
	}

	repository.UpdateOneUser = func(user *model.User, update bson.D) error {
		// emulate update success
		return nil
	}

	update := bson.D{{Key: "Password", Value: "newpass"}, {Key: "Team", Value: teamID.Hex()}}

	_, err := UpdateUserByID(uid, update)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
