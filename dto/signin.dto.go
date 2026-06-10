package dto

// SignInDTO represents the data transfer object for user sign-in operations
type SignInDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
