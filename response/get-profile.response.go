package response

// Structuring the profile response
type ProfileResponse struct {
	Status  string       `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Err     string       `json:"error,omitempty"`
	Data    UserProfiles `json:"profile,omitempty"`
}

// Structuring user profile structure
type UserProfiles struct {
	ID          int    `json:"id"`
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	UserAuthsId int    `json:"user_auths_id"`
}
