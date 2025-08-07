package models

// TestResponse represents a test response
type TestResponse struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
