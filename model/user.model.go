package model

import "gorm.io/gorm"

// User represents the user model in the database
type UserModel struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}

// UserAuth represents the user authentication model in the database
type UserAuth struct {
	gorm.Model
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Status   string `json:"status" binding:"required"`
	RolesID  uint   `json:"roles_id" binding:"required"`
}
