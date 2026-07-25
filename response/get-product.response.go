package response

import "gorm.io/gorm"

// GetProductResponse represents the response structure for retrieving products
type ProductResponse struct {
	Status  string                `json:"status"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Err     string                `json:"error,omitempty"`
	Data    []ProductResponseData `json:"data,omitempty"` // This can be a list of products or a single product
}

type ProductResponseData struct {
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	Price               float64           `json:"price"`
	Discount            float64           `json:"discount"`
	DiscountedPrice     float64           `json:"discounted_price"`
	DiscountStartDate   string            `json:"discount_start_date"`
	DiscountEndDate     string            `json:"discount_end_date"`
	Status              string            `json:"status"`
	ProductCategoriesID int               `json:"product_categories_id"`
	Logo                string            `json:"logo"`
	Business            Businesses        `json:"business"`
	Cateegories         ProductCategories `json:"category"`
}

type Products struct {
	gorm.Model
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	Price               float64           `json:"price"`
	Discount            float64           `json:"discount"`
	DiscountedPrice     float64           `json:"discounted_price"`
	Status              string            `json:"status"`
	Logo                string            `json:"logo"`
	ProductCategoriesID int               `json:"product_categories_id"`
	Categories          ProductCategories `gorm:"foreignKey:ProductCategoriesID"`
	BusinessesID        uint              `json:"businesses_id:ProductCategoriesID"`
	Businesses          Businesses        `gorm:"foreignKey:BusinessesID"`
}

type Businesses struct {
	gorm.Model
	Name        string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Description string `json:"description" binding:"required" validate:"required,min=10,max=500"`
	LogoURL     string `json:"logo_url" binding:"required"`
	Street      string `json:"street" binding:"required"`
	PopularName string `json:"popular_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Twitter     string `json:"twitter" binding:"required"`
	Facebook    string `json:"facebook" binding:"required"`
	Instagram   string `json:"instagram" binding:"required"`
	Province    string `json:"province" binding:"required"`
	District    string `json:"district" binding:"required"`
}

type ProductCategories struct {
	gorm.Model
	Name string `json:"name"`
}
