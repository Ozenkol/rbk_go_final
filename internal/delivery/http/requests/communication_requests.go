package http_requests

// CreateCommunicationRequest represents the request body for creating a communication record
// swagger:model CreateCommunicationRequest
type CreateCommunicationRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string `json:"client_id"`
}

// UpdateCommunicationRequest represents the request body for updating a communication record
// swagger:model UpdateCommunicationRequest
type UpdateCommunicationRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string `json:"client_id"`
}
