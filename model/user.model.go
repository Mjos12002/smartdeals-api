package model

import "gorm.io/gorm"

// User represents the user model in the database
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Status   string `json:"status" binding:"required"`
	RolesId  uint   `json:"roles_id" binding:"required"`
}

// UserAuth represents the user authentication model in the database
type UserAuth struct {
	gorm.Model
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Status   string `json:"status" binding:"required"`
	RolesID  uint   `json:"roles_id" binding:"required"`
}
