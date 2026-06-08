package http_requests

type CreateNotificationRequest struct {
	ClientID string `json:"client_id"`
	Message  string `json:"message"`
}

type UpdateNotificationRequest struct {
	ClientID string `json:"client_id"`
	Message  string `json:"message"`
}
