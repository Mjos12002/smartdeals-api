package dto

// ProductCategoryDTO represents the data transfer object for product categories
type ProductCategoryDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
