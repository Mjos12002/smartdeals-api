package model

import "gorm.io/gorm"

// Address represents the address model in the database
type Addresses struct {
	gorm.Model
	Street      string `json:"street"`
	PopularName string `json:"popular_name" binding:"required" validate:"required,min=5,max=50"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Sector      string `json:"sector"`
	LongLat     string `json:"long_lat"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
}
