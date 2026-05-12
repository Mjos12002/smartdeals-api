package model

import "gorm.io/gorm"

// Business represents the business details model in the database
type Business struct {
	gorm.Model
	Name          string `json:"b_name" binding:"required"`
	Description   string `json:"b_description" binding:"required"`
	Phone         string `json:"b_contact" binding:"required"`
	LogoURL       string `json:"logo_url" binding:"required"`
	AddressID     int    `json:"address_id" binding:"required"`
	UserDetailsID int    `json:"userdetails_id" binding:"required"`
}
