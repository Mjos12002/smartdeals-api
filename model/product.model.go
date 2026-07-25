package model

import "gorm.io/gorm"

// Product represents the product model in the database
type Products struct {
	gorm.Model
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Price               float64 `json:"price"`
	Discount            float64 `json:"discount"`
	DiscountedPrice     float64 `json:"discounted_price"`
	Status              string  `json:"status"`
	ProductCategoriesID int     `json:"product_categories_id"`
	Logo                string  `json:"logo"`
	BusinessesID        int     `json:"businesses_id"`
}
