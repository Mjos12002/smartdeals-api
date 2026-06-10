package utils

// This file contains validation logic for the User model
import (
	"github.com/go-playground/validator/v10"
	"smartdeals.rw/model"
)

// ValidateUserDetails validates the user details data using the validator package
func ValidateUserDetails(userDetails model.UserProfiles) error {
	validate := validator.New()
	return validate.Struct(userDetails)
}

func ValidateUserAuth(userAuth model.UserAuth) error {
	validate := validator.New()
	return validate.Struct(userAuth)
}
