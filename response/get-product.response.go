package response

import "smartdeals.rw/model"

// This file contains the response structures for the Product entity.

// GetProductResponse represents the response structure for retrieving products
type ProductResponse struct {
	Status  string           `json:"status"`
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Err     string           `json:"error,omitempty"`
	Data    []model.Products `json:"data,omitempty"` // This can be a list of products or a single product
}
