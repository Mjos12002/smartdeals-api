package model

import "gorm.io/gorm"

// UserAuth represents the user details model in the database
type UserAuths struct {
	gorm.Model
	Username      string `json:"username" binding:"required"`
	Password      string `json:"password" binding:"required,min=6"`
	Status        string `json:"status" binding:"required"`
	UserDetailsID int    `json:"user_details_id" binding:"required"`
}
