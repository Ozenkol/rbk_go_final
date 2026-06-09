package http_requests

type CreateContractRequest struct {
	ClientID   string `json:"client_id" binding:"required"`
	DealID     string `json:"deal_id"`
	Number     string `json:"number" binding:"required"`
	Status     string `json:"status"`
	ValidFrom  int64  `json:"valid_from"`
	ValidUntil int64  `json:"valid_until"`
}

type UpdateContractRequest struct {
	Number     string `json:"number"`
	Status     string `json:"status"`
	ValidFrom  int64  `json:"valid_from"`
	ValidUntil int64  `json:"valid_until"`
}
