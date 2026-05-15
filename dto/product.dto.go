package dto

// This file contains the Data Transfer Object (DTO) definitions for the Product entity.
type ProductDTO struct {
	ProductName        string  `json:"product_name" binding:"required"`
	ProductDescription string  `json:"product_description" binding:"required"`
	Price              float64 `json:"price" binding:"required"`
	Discount           float64 `json:"discount"`
	DiscountedPrice    float64 `json:"discounted_price"`
	DiscountStartDate  string  `json:"discount_start_date"`
	DiscountEndDate    string  `json:"discount_end_date"`
	ProductStatus      string  `json:"product_status" binding:"required"`
	ProductCategory    string  `json:"product_category" binding:"required"`
	ImageURL           string  `json:"image_url" binding:"required"`
	BusinessID         int     `json:"business_id" binding:"required"`
}
