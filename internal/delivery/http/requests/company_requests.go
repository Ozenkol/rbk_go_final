package http_requests

type CreateCompanyRequest struct {
	Name string `json:"name"`
}

type UpdateCompanyRequest struct {
	Name string `json:"name"`
}
