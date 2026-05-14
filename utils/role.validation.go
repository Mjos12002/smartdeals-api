package utils

// This file contains validation logic for the Roles model

import (
	"github.com/go-playground/validator/v10"
	"smartdeals.rw/model"
)

// ValidateRole validates the role data using the validator package
func ValidateRole(role model.Roles) error {
	validate := validator.New()
	return validate.Struct(role)
}
