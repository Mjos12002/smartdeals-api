package model

import "gorm.io/gorm"

// UserRole represents the user role model in the database
type UserRoleModel struct {
	gorm.Model
	UserAuthID UserAuthModel `json:"userauth_id" binding:"required"`
	RoleID     Roles         `json:"roles_id" binding:"required"`
}
