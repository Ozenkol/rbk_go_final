package http_requests

// CreateFileRequest represents the request body for uploading a file
// swagger:model CreateFileRequest
type CreateFileRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID         string `json:"client_id"`
	// Storage reference path
	// example: uploads/2023/10/file.pdf
	StorageReference string `json:"storage_reference"`
	// File name
	// example: contract.pdf
	Name             string `json:"name"`
}

// UpdateFileRequest represents the request body for updating file metadata
// swagger:model UpdateFileRequest
type UpdateFileRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID         string `json:"client_id"`
	// Storage reference path
	// example: uploads/2023/10/file.pdf
	StorageReference string `json:"storage_reference"`
	// File name
	// example: contract_v2.pdf
	Name             string `json:"name"`
}
