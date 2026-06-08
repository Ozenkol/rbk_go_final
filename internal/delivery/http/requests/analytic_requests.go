package http_requests

type CreateAnalyticRequest struct {
	ClientID string      `json:"client_id"`
	Data     interface{} `json:"data"`
}

type UpdateAnalyticRequest struct {
	ClientID string      `json:"client_id"`
	Data     interface{} `json:"data"`
}
