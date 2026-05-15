package utils

// This file contains validation functions for the Product entity.

import (
	"github.com/go-playground/validator/v10"
	"smartdeals.rw/dto"
)

// ValidateProductDTO validates the ProductDTO struct using the validator package
func ValidateProduct(productDTO *dto.ProductDTO) error {
	validate := validator.New()
	return validate.Struct(productDTO)
}
