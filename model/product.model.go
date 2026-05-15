package model

import "gorm.io/gorm"

// Product represents the product model in the database
type Products struct {
	gorm.Model
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	Price              float64 `json:"price"`
	Discount           float64 `json:"discount"`
	DiscountedPrice    float64 `json:"discounted_price"`
	DiscountStartDate  string  `json:"discount_start_date"`
	DiscountEndDate    string  `json:"discount_end_date"`
	ProductStatus      string  `json:"product_status"`
	ProductCategory    string  `json:"product_category"`
	ImageURL           string  `json:"image_url"`
	BusinessID         int     `json:"business_id"`
}
