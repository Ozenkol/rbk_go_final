package http_requests

type CreateOfferRequest struct {
	ClientID   string `json:"client_id"`
	DocumentID string `json:"document_id"`
	CompanyID  string `json:"company_id"`
}

type UpdateOfferRequest struct {
	ClientID   string `json:"client_id"`
	DocumentID string `json:"document_id"`
	CompanyID  string `json:"company_id"`
}
