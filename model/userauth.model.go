package model

import "gorm.io/gorm"

// UserAuth represents the user details model in the database
type UserAuth struct {
	gorm.Model
	Username    string           `json:"username" binding:"required"`
	Password    string           `json:"u_password" binding:"required,min=6"`
	Status      string           `json:"status" binding:"required"`
	UserDetails UserDetailsModel `json:"user_details" binding:"required"`
}
