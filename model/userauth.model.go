package model

// UserAuth represents the user details model in the database
type UserAuth struct {
	Username    string           `json:"username" binding:"required"`
	UserDetails UserDetailsModel `json:"user_details" binding:"required"`
}
