package response

// GenericCreateResponse represents the response structure for a successful creation operation
type GenericCreateResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Err      string `json:"error"`
	RecordID int    `json:"record_id"`
	Token    string `json:"token"`
	UserID   int    `json:"id"`
	UserRole string `json:"role"`
	Username string `json:"username"`
}
