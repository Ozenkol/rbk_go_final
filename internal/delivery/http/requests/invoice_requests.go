package http_requests

type CreateInvoiceRequest struct {
	ClientID   string `json:"client_id"`
	DocumentID string `json:"document_id"`
	CompanyID  string `json:"company_id"`
}

type UpdateInvoiceRequest struct {
	ClientID   string `json:"client_id"`
	DocumentID string `json:"document_id"`
	CompanyID  string `json:"company_id"`
}
