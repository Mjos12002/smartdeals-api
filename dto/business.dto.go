package dto

import "gorm.io/gorm"

// BusinessesDTO represents the data transfer object for business details
type BusinessesDTO struct {
	gorm.Model
	Name        string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Description string `json:"description" binding:"required" validate:"required,min=10,max=500"`
	Contact     string `json:"contact" binding:"required" validate:"required,min=10,max=15"`
	Phone       string `json:"phone" binding:"required" validate:"required,min=10,max=15"`
	LogoURL     string `json:"logo_url" binding:"required"`
	Address     int    `json:"address_id" binding:"required" validate:"required"`
	UserDetails int    `json:"user_details_id" binding:"required" validate:"required"`
}
