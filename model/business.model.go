package model

import "gorm.io/gorm"

// Business represents the business details model in the database
type BusinessModel struct {
	gorm.Model
	Name        string      `json:"b_name" binding:"required"`
	Description string      `json:"b_description" binding:"required"`
	Phone       string      `json:"b_contact" binding:"required"`
	LogoURL     string      `json:"logo_url" binding:"required"`
	Address     Addresses   `json:"address_id" binding:"required"`
	UserDetails UserDetails `json:"address" binding:"required"`
}
