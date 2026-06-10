package response

import "smartdeals.rw/model"

// GetUserResponse represents the response structure for user-related operations
type UserResponse struct {
	Status  string               `json:"status"`
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Err     string               `json:"error,omitempty"`
	Data    []model.UserProfiles `json:"data,omitempty"`
}
