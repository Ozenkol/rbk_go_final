package http_requests

type CreateSettingRequest struct {
	CompanyID string `json:"company_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type UpdateSettingRequest struct {
	CompanyID string `json:"company_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
