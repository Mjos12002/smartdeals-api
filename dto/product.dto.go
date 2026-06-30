package dto

// ProductDTOProductDTO structures the data transfer object (DTO) for the product .
type ProductDTO struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description" binding:"required"`
	Price             int    `json:"price" binding:"required"`
	Discount          int    `json:"discount"`
	DiscountedPrice   int    `json:"discounted_price"`
	DiscountStartDate string `json:"discount_start_date"`
	DiscountEndDate   string `json:"discount_end_date"`
	Status            string `json:"status" binding:"required"`
	ProductCategory   int    `json:"product_categories_id" binding:"required"`
	Logo              string `json:"logo" binding:"required"`
	BusinessID        int    `json:"business_id" binding:"required"`
}
