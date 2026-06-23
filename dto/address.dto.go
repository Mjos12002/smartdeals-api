package dto

// This file contains the Data Transfer Object (DTO) definitions for the Address entity.
type AddressDTO struct {
	Street      string `json:"street" binding:"required,min=5,max=100"`
	PopularName string `json:"popular_name" binding:"required,min=5,max=100"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Sector      string `json:"sector"`
	LongLat     string `json:"long_lat"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
}
