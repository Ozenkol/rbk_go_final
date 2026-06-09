package http_requests

type CreateInvoiceRequest struct {
	ClientID  string `json:"client_id" binding:"required"`
	DealID    string `json:"deal_id"`
	Number    string `json:"number" binding:"required"`
	DueDate   int64  `json:"due_date"`
}

type UpdateInvoiceRequest struct {
	Status     string  `json:"status"`
	PaidAmount float64 `json:"paid_amount"`
}
