package dto

type SignUpDTO struct {
	Username string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
