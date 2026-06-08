package http_requests

// CreateInvoiceRequest represents the request body for creating an invoice
// swagger:model CreateInvoiceRequest
type CreateInvoiceRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID   string `json:"client_id"`
	// Document ID
	// example: 550e8400-e29b-41d4-a716-446655440001
	DocumentID string `json:"document_id"`
}

// UpdateInvoiceRequest represents the request body for updating an invoice
// swagger:model UpdateInvoiceRequest
type UpdateInvoiceRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID   string `json:"client_id"`
	// Document ID
	// example: 550e8400-e29b-41d4-a716-446655440001
	DocumentID string `json:"document_id"`
}
