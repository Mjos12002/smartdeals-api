package model

import "gorm.io/gorm"

// ProductCategory represents the product category model in the database
type ProductCategory struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
