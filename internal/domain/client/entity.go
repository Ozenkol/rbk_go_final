package client

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Client struct {
	ID                string            `json:"id"`
	// swagger:ignore
	UserID            string            `json:"user_id"`
	// swagger:ignore
	CompanyID         string            `json:"company_id"`
	Type              shared.ClientType `json:"type"` // Physical/Legal
	Name              string            `json:"name"` // Company name or Full name
	Email             string            `json:"email"`
	Phone             string            `json:"phone"`
	WhatsApp          string            `json:"whatsapp"`
	IdentificationNum string            `json:"identification_num"` // BIN/IIN
	Address           string            `json:"address"`
	Source            string            `json:"source"`
	Status            string            `json:"status"`
	ResponsibleID     string            `json:"responsible_id"`
	Tags              []string          `json:"tags"`
	Person            shared.Person     `json:"person"`
	IsActive          bool              `json:"is_active"`
	CreatedAt         int64             `json:"created_at"`
	LastContactAt     int64             `json:"last_contact_at"`
	Comment           string            `json:"comment"`
}

func NewClient(clientType shared.ClientType, name string, isActive bool) Client {
	return Client{
		Type:     clientType,
		Name:     name,
		IsActive: isActive,
		Tags:     []string{},
	}
}
