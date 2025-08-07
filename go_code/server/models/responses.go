package models

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message"`
}

// ValueResponse represents a value response
type ValueResponse struct {
	Value interface{} `json:"value"`
}
