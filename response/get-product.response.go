package response

// GetProductResponse represents the response structure for retrieving products
type ProductResponse struct {
	Status  string                `json:"status"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Err     string                `json:"error,omitempty"`
	Data    []ProductResponseData `json:"data,omitempty"` // This can be a list of products or a single product
}

type ProductResponseData struct {
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Price               float64 `json:"price"`
	Discount            float64 `json:"discount"`
	DiscountedPrice     float64 `json:"discounted_price"`
	DiscountStartDate   string  `json:"discount_start_date"`
	DiscountEndDate     string  `json:"discount_end_date"`
	Status              string  `json:"status"`
	ProductCategoriesID int     `json:"product_categories_id"`
	Logo                string  `json:"logo"`
	BusinessesID        int     `json:"businesses_id"`
}
