package teams

import (
	"testing"
)

func TestTeamOutput_Struct(t *testing.T) {
	output := TeamOutput{
		ID:   "507f1f77bcf86cd799439011",
		Name: "Engineering",
	}

	if output.ID != "507f1f77bcf86cd799439011" {
		t.Fatalf("expected ID to match")
	}

	if output.Name != "Engineering" {
		t.Fatalf("expected Name to be Engineering, got %s", output.Name)
	}
}

func TestTeamOutput_DifferentName(t *testing.T) {
	output := TeamOutput{
		ID:   "507f1f77bcf86cd799439012",
		Name: "Product",
	}

	if output.Name != "Product" {
		t.Fatalf("expected Name to be Product, got %s", output.Name)
	}
}

func TestTeamOutput_MultipleTeams(t *testing.T) {
	teamNames := []string{"Engineering", "Product", "Design", "Marketing"}
	for i, name := range teamNames {
		output := TeamOutput{
			ID:   "507f1f77bcf86cd79943901" + string(rune(i+1)),
			Name: name,
		}
		if output.Name != name {
			t.Fatalf("expected Name to be %s, got %s", name, output.Name)
		}
	}
}
