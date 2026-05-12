package model

import "gorm.io/gorm"

// UserRole represents the user role model in the database
type UserRole struct {
	gorm.Model
	UserAuthID UserAuth `json:"userauth_id" binding:"required"`
	RoleID     Role     `json:"roles_id" binding:"required"`
}
