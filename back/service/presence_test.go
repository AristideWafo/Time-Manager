package service

import (
	"errors"
	"testing"
	"time"

	"TimeManager/model"
	"TimeManager/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCreatePresence_Success(t *testing.T) {
	orig := repository.SavePresence
	defer func() { repository.SavePresence = orig }()

	var saved *model.Presence
	repository.SavePresence = func(p *model.Presence) error {
		saved = p
		return nil
	}

	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	ts := time.Now()

	p, err := CreatePresence("IN", userID, ts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if p == nil {
		t.Fatalf("expected presence, got nil")
	}
	if saved == nil {
		t.Fatalf("expected SavePresence to be called")
	}
	if saved.User != userID || saved.Type != "IN" || !saved.Timestamp.Equal(ts) {
		t.Fatalf("saved presence fields mismatch: got %+v", saved)
	}
}

func TestCreatePresence_SaveError(t *testing.T) {
	orig := repository.SavePresence
	defer func() { repository.SavePresence = orig }()

	repository.SavePresence = func(p *model.Presence) error {
		return errors.New("db fail")
	}

	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439011")
	ts := time.Now()

	p, err := CreatePresence("OUT", userID, ts)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if p == nil {
		t.Fatalf("expected presence to still be returned even on save error")
	}
}

func TestGetAllPresencesByUserID_Success(t *testing.T) {
	orig := repository.GetManyPresences
	defer func() { repository.GetManyPresences = orig }()

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		uid := bson.ObjectID{}
		*presences = []*model.Presence{
			{Type: "IN", Timestamp: time.Unix(1, 0), User: uid},
		}
		return nil
	}

	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	ps, err := GetAllPresencesByUserID(userID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ps == nil || len(*ps) != 1 {
		t.Fatalf("expected 1 presence, got %v", ps)
	}
}

func TestGetAllPresencesByUserID_Error(t *testing.T) {
	orig := repository.GetManyPresences
	defer func() { repository.GetManyPresences = orig }()

	repository.GetManyPresences = func(presences *[]*model.Presence, filter bson.D) error {
		return errors.New("db error")
	}

	userID, _ := bson.ObjectIDFromHex("507f1f77bcf86cd799439012")
	_, err := GetAllPresencesByUserID(userID)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
