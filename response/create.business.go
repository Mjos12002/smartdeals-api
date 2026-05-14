package response

import "smartdeals.rw/model"

// Get BusinessResponse represents the response structure for business-related operations
type BusinessResponse struct {
	Status  string             `json:"status"`
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Err     string             `json:"error,omitempty"`
	Data    []model.Businesses `json:"data,omitempty"`
}
