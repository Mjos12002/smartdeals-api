package model

// UserDetails represents the user details model in the database
type UserDetailsModel struct {
	Email   string  `json:"email" binding:"required,email"`
	Phone   string  `json:"phone" binding:"required"`
	Address Address `json:"u_address" binding:"required"`
	FName   string  `json:"fname" binding:"required"`
	LName   string  `json:"lname" binding:"required"`
}
