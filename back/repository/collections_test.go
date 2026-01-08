package repository

import (
	"testing"

	"github.com/subosito/gotenv"
)

func setupCollectionsTest(t *testing.T) {
	err := gotenv.Load("../.env")
	if err != nil {
		t.Fatalf("failed to load .env file: %v", err)
	}
	DBConnect()
	if Client.Database == nil {
		t.Skip("Database not connected")
	}
}

func TestTeamCollection_InitializesOnce(t *testing.T) {
	setupCollectionsTest(t)

	col1 := TeamCollection()
	col2 := TeamCollection()

	if col1 != col2 {
		t.Fatalf("expected same collection instance on multiple calls")
	}
}

func TestUserCollection_InitializesOnce(t *testing.T) {
	setupCollectionsTest(t)

	col1 := UserCollection()
	col2 := UserCollection()

	if col1 != col2 {
		t.Fatalf("expected same collection instance on multiple calls")
	}
}

func TestPresenceCollection_InitializesOnce(t *testing.T) {
	setupCollectionsTest(t)

	col1 := PresenceCollection()
	col2 := PresenceCollection()

	if col1 != col2 {
		t.Fatalf("expected same collection instance on multiple calls")
	}
}
