package dto

// This file contains the Data Transfer Object (DTO) definitions for the Address entity.
type AddressDTO struct {
	Street      string `json:"street" binding:"required,min=5,max=100"`
	PopularName string `json:"popular_name" binding:"required,min=5,max=100"`
	Province    string `json:"province" binding:"required,min=5,max=100"`
	District    string `json:"district" binding:"required"`
	Sector      string `json:"sector" binding:"required"`
	LongLat     string `json:"longlat" binding:"required"`
}
