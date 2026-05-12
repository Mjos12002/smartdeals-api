package model

import "gorm.io/gorm"

// Address represents the address model in the database
type Addresses struct {
	gorm.Model
	Street      string `json:"street"`
	PopularName string `json:"popularname"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Sector      string `json:"sector"`
	LongLat     string `json:"longlat"`
}
