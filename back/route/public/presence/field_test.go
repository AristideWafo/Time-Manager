package presence

import (
	"testing"
	"time"
)

func TestPresenceOutput_Struct(t *testing.T) {
	now := time.Now()
	output := PresenceOutput{
		Type:      "PRESENT",
		Timestamp: now,
	}

	if output.Type != "PRESENT" {
		t.Fatalf("expected Type to be PRESENT, got %s", output.Type)
	}

	if output.Timestamp != now {
		t.Fatalf("expected Timestamp to match")
	}
}

func TestPresenceOutput_DifferentType(t *testing.T) {
	output := PresenceOutput{
		Type:      "ABSENT",
		Timestamp: time.Now(),
	}

	if output.Type != "ABSENT" {
		t.Fatalf("expected Type to be ABSENT, got %s", output.Type)
	}
}

func TestCreatePresenceInput_Struct(t *testing.T) {
	now := time.Now()
	input := CreatePresenceInput{
		Type:      "PRESENT",
		Timestamp: now,
	}

	if input.Type != "PRESENT" {
		t.Fatalf("expected Type to be PRESENT, got %s", input.Type)
	}

	if input.Timestamp != now {
		t.Fatalf("expected Timestamp to match")
	}
}

func TestCreatePresenceInput_DifferentTypes(t *testing.T) {
	types := []string{"PRESENT", "ABSENT", "LATE", "EARLY_LEAVE"}
	for _, presenceType := range types {
		input := CreatePresenceInput{
			Type:      presenceType,
			Timestamp: time.Now(),
		}
		if input.Type != presenceType {
			t.Fatalf("expected Type to be %s, got %s", presenceType, input.Type)
		}
	}
}
