package model

import "gorm.io/gorm"

// Role represents the role model in the database
type Role struct {
	gorm.Model
	RoleName string `json:"r_name"`
}
