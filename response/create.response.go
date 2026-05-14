package response

// CreateResponse represents the response structure for a successful creation operation
type CreateResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Err      string `json:"error,omitempty"`
	RecordID int    `json:"record_id,omitempty"`
}
