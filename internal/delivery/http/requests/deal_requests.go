package http_requests

type CreateDealRequest struct {
	ClientID      string  `json:"client_id" binding:"required"`
	ResponsibleID string  `json:"responsible_id"`
	Title         string  `json:"title" binding:"required"`
	Stage         string  `json:"stage"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Probability   int     `json:"probability"`
	Deadline      int64   `json:"deadline"`
}

type UpdateDealRequest struct {
	ResponsibleID string  `json:"responsible_id"`
	Title         string  `json:"title"`
	Stage         string  `json:"stage"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
	Probability   int     `json:"probability"`
	Deadline      int64   `json:"deadline"`
}
