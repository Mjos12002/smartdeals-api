package response

// This file contains the response structures for the Business entity.

import "smartdeals.rw/model"

// BusinessResponse represents the response structure for retrieving businesses
type BusinessResponse struct {
	Status  string             `json:"status"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Err     string             `json:"error,omitempty"`
	Data    []model.Businesses `json:"data,omitempty"` // This can be a list of businesses or a single business
}
