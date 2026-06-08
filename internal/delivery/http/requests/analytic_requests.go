package http_requests

// CreateAnalyticRequest represents the request body for creating an analytic record
// swagger:model CreateAnalyticRequest
type CreateAnalyticRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string      `json:"client_id"`
	// Analytic data
	Data     interface{} `json:"data"`
}

// UpdateAnalyticRequest represents the request body for updating an analytic record
// swagger:model UpdateAnalyticRequest
type UpdateAnalyticRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string      `json:"client_id"`
	// Analytic data
	Data     interface{} `json:"data"`
}
