package http_requests

type CreateDocumentRequest struct {
	ClientID       string `json:"client_id"`
	CompanyID      string `json:"company_id"`
	Type           string `json:"type"`
	Number         string `json:"number"`
	IssuedBy       string `json:"issued_by"`
	IssuedDate     string `json:"issued_date"`
	ExpirationDate string `json:"expiration_date"`
}

type UpdateDocumentRequest struct {
	ClientID       string `json:"client_id"`
	CompanyID      string `json:"company_id"`
	Type           string `json:"type"`
	Number         string `json:"number"`
	IssuedBy       string `json:"issued_by"`
	IssuedDate     string `json:"issued_date"`
	ExpirationDate string `json:"expiration_date"`
}
