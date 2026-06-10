package model

// SignInResponseModel represents the data structure for user sign-in response
type SignInResponseModel struct {
	Status string `json:"status"`
	ID     uint   `json:"id,omitempty"`
}
