package model

import "gorm.io/gorm"

// UserDetails represents the user details model in the database
type UserProfiles struct {
	gorm.Model
	Email     string `json:"email" binding:"required" validate:"required,email,min=5,max=50"`
	Phone     string `json:"phone" binding:"required" validate:"required,min=10,max=15"`
	FirstName string `json:"first_name" binding:"required" validate:"required,min=5,max=30"`
	LastName  string `json:"last_name" binding:"required" validate:"required,min=5,max=30"`
}
