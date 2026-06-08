package http_requests

type CreateCommunicationRequest struct {
	ClientID string `json:"client_id"`
}

type UpdateCommunicationRequest struct {
	ClientID string `json:"client_id"`
}
