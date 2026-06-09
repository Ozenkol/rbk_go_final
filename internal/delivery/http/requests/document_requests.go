package http_requests

type CreateDocumentRequest struct {
	ClientID       string `json:"client_id"`
	DealID         string `json:"deal_id"`
	ContractID     string `json:"contract_id"`
	Type           string `json:"type" binding:"required"`
	Number         string `json:"number"`
	IssuedBy       string `json:"issued_by"`
	IssuedDate     int64  `json:"issued_date"`
	ExpirationDate int64  `json:"expiration_date"`
}

type UpdateDocumentRequest struct {
	Type           string `json:"type"`
	Number         string `json:"number"`
	IssuedBy       string `json:"issued_by"`
	IssuedDate     int64  `json:"issued_date"`
	ExpirationDate int64  `json:"expiration_date"`
}
