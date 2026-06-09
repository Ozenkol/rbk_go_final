package http_requests

type CreateProposalRequest struct {
	ClientID string  `json:"client_id" binding:"required"`
	DealID   string  `json:"deal_id"`
	Title    string  `json:"title" binding:"required"`
	Currency string  `json:"currency"`
	Discount float64 `json:"discount"`
}

type UpdateProposalRequest struct {
	Title    string  `json:"title"`
	Status   string  `json:"status"`
	Discount float64 `json:"discount"`
}
