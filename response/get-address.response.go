package response

import "smartdeals.rw/model"

// GetAddressResponse represents the response structure for address-related operations
type GetAddressResponse struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Err     string            `json:"error,omitempty"`
	Data    []model.Addresses `json:"data,omitempty"`
}
