package dto

// This file contains the Data Transfer Object (DTO) definitions for the Product entity.
type ProductDTO struct {
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
	Discount          float64 `json:"discount"`
	DiscountedPrice   float64 `json:"discounted_price"`
	DiscountStartDate string  `json:"discount_start_date"`
	DiscountEndDate   string  `json:"discount_end_date"`
	Status            string  `json:"status" binding:"required"`
	ProductCategory   int     `json:"product_category" binding:"required"`
	Logo              string  `json:"logo" binding:"required"`
	BusinessID        int     `json:"business_id" binding:"required"`
}
