package http_requests

type CreateContractRequest struct {
	ClientID  string `json:"client_id"`
	CompanyID string `json:"company_id"`
	Content   string `json:"content"`
}

type UpdateContractRequest struct {
	ClientID  string `json:"client_id"`
	CompanyID string `json:"company_id"`
	Content   string `json:"content"`
}
