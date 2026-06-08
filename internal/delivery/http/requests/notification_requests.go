package http_requests

// CreateNotificationRequest represents the request body for creating a notification
// swagger:model CreateNotificationRequest
type CreateNotificationRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string `json:"client_id"`
	// Notification message
	// example: Your subscription is about to expire
	Message  string `json:"message"`
}

// UpdateNotificationRequest represents the request body for updating a notification
// swagger:model UpdateNotificationRequest
type UpdateNotificationRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID string `json:"client_id"`
	// Notification message
	// example: Your subscription has expired
	Message  string `json:"message"`
}
