package utils

// This file contains the validation functions for the Product Category data.

import (
	"github.com/go-playground/validator/v10"
	"smartdeals.rw/dto"
)

// ValidateProductCategory validates the incoming product category data
func ValidateProductCategory(categoryDTO *dto.ProductCategoryDTO) error {
	validate := validator.New()
	return validate.Struct(categoryDTO)
}
