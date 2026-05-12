package model

// Product represents the product model in the database
type Product struct {
	ID                int      `json:"id"`
	Name              string   `json:"product_name"`
	Description       string   `json:"product_description"`
	Price             float64  `json:"price"`
	Discount          float64  `json:"discount"`
	DiscountedPrice   float64  `json:"discounted_price"`
	DiscountStartDate string   `json:"discount_start_date"`
	DiscountEndDate   string   `json:"discount_end_date"`
	Status            string   `json:"product_status"`
	Category          string   `json:"product_category"`
	ImageURL          string   `json:"image_url"`
	BusinessID        Business `json:"business_id"`
}
