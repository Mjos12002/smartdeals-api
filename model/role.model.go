package model

import "gorm.io/gorm"

// Role represents the role model in the database
type Roles struct {
	gorm.Model
	RoleName string `json:"role_name" binding:"required"`
}
