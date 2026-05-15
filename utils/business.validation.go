package utils

import (
	"github.com/go-playground/validator/v10"
	"smartdeals.rw/dto"
)

// This file contains validation logic for the Business model

// ValidateBusiness validates the business data using the validator package
func ValidateBusiness(business dto.BusinessesDTO) error {
	validate := validator.New()
	return validate.Struct(business)
}
