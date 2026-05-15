package response

import "smartdeals.rw/model"

// This file contains the response structures for the Product Category responses.

// GetProductCategoryResponse represents the response structure for fetching product categories
type ProductCategoryResponse struct {
	Status  string                  `json:"status"`
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Err     string                  `json:"err"`
	Data    []model.ProductCategory `json:"data"`
}
