package model

import "github.com/go-playground/validator/v10"

func ValidateModel(structure interface{}) error {

	validate := validator.New()
	return validate.Struct(structure)
}
