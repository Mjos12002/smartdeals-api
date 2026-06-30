package model

import "gorm.io/gorm"

// Business represents the business details model in the database
type Businesses struct {
	gorm.Model
	Name           string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Description    string `json:"description" binding:"required" validate:"required,min=10,max=500"`
	LogoURL        string `json:"logo_url" binding:"required"`
	AddressesID    int    `json:"addresses_id" binding:"required" validate:"required"`
	UserProfilesID int    `json:"user_profiles_id" binding:"required" validate:"required"`
}
