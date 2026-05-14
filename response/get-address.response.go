package response

import "smartdeals.rw/model"

// AddressResponse represents the response structure for address-related operations
type AddressResponse struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Err     string            `json:"error,omitempty"`
	Data    []model.Addresses `json:"data,omitempty"`
}
