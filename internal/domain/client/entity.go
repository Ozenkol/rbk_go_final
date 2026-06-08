package client

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Client struct {
	ID        string `json:"id"`
	// swagger:ignore
	UserID    string `json:"user_id"`
	// swagger:ignore
	CompanyID string `json:"company_id"`
	Person    shared.Person `json:"person"`
	IsActive  bool `json:"is_active"`
}

func NewClient(person shared.Person, isActive bool) Client {
	return Client{
		Person:   person,
		IsActive: isActive,
	}
}
