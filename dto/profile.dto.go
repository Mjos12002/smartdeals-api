package dto

// ProfileDTO is a structure of the user profile.
type ProfileDTO struct {
	FirstName   string `json:"first_name" binding:"required,min=5,max=30"`
	LastName    string `json:"last_name" binding:"required,min=5,max=30"`
	Email       string `json:"email" binding:"required,min=5,max=30"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
