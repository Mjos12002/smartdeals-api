package response

// This file contains the response structures for the Product entity.

// CreateProductResponse represents the response structure for creating a product
type CreateProductResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Err      string `json:"error,omitempty"`
	RecordID int    `json:"record_id"` // Assuming the created product's ID is returned
}
