package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

// CreateClientRequest represents the request body for creating a client
// swagger:model CreateClientRequest
type CreateClientRequest struct {
	// Client personal information
	Person    shared.Person `json:"person"`
	// Client active status
	// example: true
	IsActive  bool          `json:"is_active"`
}

// UpdateClientRequest represents the request body for updating a client
// swagger:model UpdateClientRequest
type UpdateClientRequest struct {
	// Client personal information
	Person    shared.Person `json:"person"`
	// Client active status
	// example: true
	IsActive  bool          `json:"is_active"`
}
