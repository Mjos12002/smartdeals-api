package model

import "gorm.io/gorm"

// Product represents the product model in the database
type Products struct {
	gorm.Model
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Discount          float64 `json:"discount"`
	DiscountedPrice   float64 `json:"discounted_price"`
	DiscountStartDate string  `json:"discount_start_date"`
	DiscountEndDate   string  `json:"discount_end_date"`
	Status            string  `json:"status"`
	ProductCategory   int     `json:"product_category"`
	Logo              string  `json:"logo"`
	BusinessID        int     `json:"business_id"`
}
