package model

import "gorm.io/gorm"

// UserDetails represents the user details model in the database
type UserProfiles struct {
	gorm.Model
	FirstName   string `json:"first_name" binding:"required" validate:"required,min=5,max=30"`
	LastName    string `json:"last_name" binding:"required" validate:"required,min=5,max=30"`
	UserAuthsId int    `json:"user_auths_id"`
}
