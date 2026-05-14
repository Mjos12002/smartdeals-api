package response

import "smartdeals.rw/model"

// RoleResponse represents the response structure for role-related operations
type RoleResponse struct {
	Status  string        `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Err     string        `json:"error,omitempty"`
	Data    []model.Roles `json:"data,omitempty"`
}
