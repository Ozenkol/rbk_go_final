package http_requests

// CreateSettingRequest represents the request body for creating a setting
// swagger:model CreateSettingRequest
type CreateSettingRequest struct {
	// Setting key
	// example: maintenance_mode
	Key       string `json:"key"`
	// Setting value
	// example: false
	Value     string `json:"value"`
}

// UpdateSettingRequest represents the request body for updating a setting
// swagger:model UpdateSettingRequest
type UpdateSettingRequest struct {
	// Setting key
	// example: maintenance_mode
	Key       string `json:"key"`
	// Setting value
	// example: true
	Value     string `json:"value"`
}
