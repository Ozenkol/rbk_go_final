package http_requests

// CreateContractRequest represents the request body for creating a contract
// swagger:model CreateContractRequest
type CreateContractRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID  string `json:"client_id"`
	// Contract content
	// example: Standard service agreement
	Content   string `json:"content"`
}

// UpdateContractRequest represents the request body for updating a contract
// swagger:model UpdateContractRequest
type UpdateContractRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID  string `json:"client_id"`
	// Contract content
	// example: Updated service agreement
	Content   string `json:"content"`
}
