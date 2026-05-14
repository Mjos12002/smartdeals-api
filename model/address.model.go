package model

import "gorm.io/gorm"

// Address represents the address model in the database
type Addresses struct {
	gorm.Model
	Street      string `json:"street" binding:"required" validate:"required,min=5,max=100"`
	PopularName string `json:"popular_name" binding:"required" validate:"required,min=5,max=50"`
	Province    string `json:"province" binding:"required" validate:"required,min=5,max=50"`
	District    string `json:"district" binding:"required"`
	Sector      string `json:"sector" binding:"required"`
	LongLat     string `json:"longlat" binding:"required"`
}
