package model

import (
	"testing"
)

type TestStruct struct {
	ID       string `validate:"required"`
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"omitempty,min=0,max=150"`
	Optional string
}

func TestValidateModel_ValidStruct(t *testing.T) {
	s := TestStruct{
		ID:    "123",
		Name:  "Doctor Who",
		Email: "Doctor@Tardis.com",
		Age:   30,
	}

	err := ValidateModel(s)
	if err != nil {
		t.Fatalf("expected no error for valid struct, got %v", err)
	}
}

func TestValidateModel_MissingRequiredField(t *testing.T) {
	s := TestStruct{
		ID:   "123",
		Name: "Doctor Who",
	}

	err := ValidateModel(s)
	if err == nil {
		t.Fatalf("expected error for missing required field, got nil")
	}
}

func TestValidateModel_InvalidEmail(t *testing.T) {
	s := TestStruct{
		ID:    "123",
		Name:  "Doctor Who",
		Email: "invalid-email",
		Age:   30,
	}

	err := ValidateModel(s)
	if err == nil {
		t.Fatalf("expected error for invalid email format, got nil")
	}
}

func TestValidateModel_OutOfRangeValue(t *testing.T) {
	s := TestStruct{
		ID:    "123",
		Name:  "Doctor Who",
		Email: "Doctor@Tardis.com",
		Age:   200,
	}

	err := ValidateModel(s)
	if err == nil {
		t.Fatalf("expected error for out-of-range age, got nil")
	}
}

func TestValidateModel_EmptyStruct(t *testing.T) {
	s := TestStruct{}

	err := ValidateModel(s)
	if err == nil {
		t.Fatalf("expected error for empty struct (missing required fields), got nil")
	}
}

func TestValidatePreSave_ValidStruct(t *testing.T) {
	s := TestStruct{
		Name:  "Doctor Who",
		Email: "Doctor@Tardis.com",
		Age:   30,
	}

	err := ValidatePreSave(s)
	if err != nil {
		t.Fatalf("expected no error for valid struct (pre-save), got %v", err)
	}
}

func TestValidatePreSave_IgnoresID(t *testing.T) {
	s := TestStruct{
		ID:    "",
		Name:  "Doctor Who",
		Email: "Doctor@Tardis.com",
		Age:   30,
	}

	err := ValidatePreSave(s)
	if err != nil {
		t.Fatalf("expected no error (ID should be ignored), got %v", err)
	}
}

func TestValidatePreSave_MissingNameField(t *testing.T) {
	s := TestStruct{
		Email: "Doctor@Tardis.com",
		Age:   30,
	}

	err := ValidatePreSave(s)
	if err == nil {
		t.Fatalf("expected error for missing Name field, got nil")
	}
}

func TestValidatePreSave_MissingEmailField(t *testing.T) {
	s := TestStruct{
		Name: "Doctor Who",
		Age:  30,
	}

	err := ValidatePreSave(s)
	if err == nil {
		t.Fatalf("expected error for missing Email field, got nil")
	}
}

func TestValidatePreSave_InvalidEmail(t *testing.T) {
	s := TestStruct{
		Name:  "Doctor Who",
		Email: "not-an-email",
		Age:   30,
	}

	err := ValidatePreSave(s)
	if err == nil {
		t.Fatalf("expected error for invalid email, got nil")
	}
}

func TestValidatePreSave_ValidEmptyOptionalField(t *testing.T) {
	s := TestStruct{
		Name:     "Doctor Who",
		Email:    "Doctor@Tardis.com",
		Age:      0,
		Optional: "",
	}

	err := ValidatePreSave(s)
	if err != nil {
		t.Fatalf("expected no error for valid struct with empty optional fields, got %v", err)
	}
}
