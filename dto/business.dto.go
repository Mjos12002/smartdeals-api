package dto

import "gorm.io/gorm"

// BusinessesDTO represents the data transfer object for business details
type BusinessesDTO struct {
	gorm.Model
	Name        string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Description string `json:"description" binding:"required" validate:"required,min=10,max=500"`
	LogoURL     string `json:"logo_url" binding:"required"`
	Street      string `json:"street" binding:"required"`
	PopularName string `json:"popular_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Twitter     string `json:"twitter" binding:"required"`
	Facebook    string `json:"facebook" binding:"required"`
	Instagram   string `json:"instagram" binding:"required"`
	Province    string `json:"province" binding:"required"`
	District    string `json:"district" binding:"required"`

	UserProfile int `json:"user_profile_id" binding:"required" validate:"required"`
}
