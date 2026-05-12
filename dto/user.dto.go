package dto

// UserDTO represents the data transfer object for user-related operations
type UserDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}
