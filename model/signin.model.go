package model

// SignInModel represents the data structure for user sign-in
type SignInModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
