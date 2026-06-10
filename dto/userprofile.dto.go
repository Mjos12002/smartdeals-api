package dto

// This file contains the Data Transfer Object (DTO) definitions for the UserDetails entity.
type UserProfileDTO struct {
	Email     string `json:"email" binding:"required,min=10,max=15"`
	Phone     string `json:"phone" binding:"required,min=10,max=15"`
	FirstName string `json:"first_name" binding:"required,min=5,max=30"`
	LastName  string `json:"last_name" binding:"required,min=5,max=30"`
	Username  string `json:"username" binding:"required,min=4,max=20"`
	
}
