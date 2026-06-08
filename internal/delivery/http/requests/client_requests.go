package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type CreateClientRequest struct {
	CompanyID string        `json:"company_id"`
	Person    shared.Person `json:"person"`
	IsActive  bool          `json:"is_active"`
}

type UpdateClientRequest struct {
	CompanyID string        `json:"company_id"`
	Person    shared.Person `json:"person"`
	IsActive  bool          `json:"is_active"`
}
