package model

import "gorm.io/gorm"

// UserDetails represents the user details model in the database
type UserDetailsModel struct {
	gorm.Model
	Email   string  `json:"email" binding:"required,email"`
	Phone   string  `json:"phone" binding:"required"`
	Address Address `json:"u_address" binding:"required"`
	FName   string  `json:"fname" binding:"required"`
	LName   string  `json:"lname" binding:"required"`
}
