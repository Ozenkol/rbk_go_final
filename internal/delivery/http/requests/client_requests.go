package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type CreateClientRequest struct {
	Type              string        `json:"type" binding:"required"`
	Name              string        `json:"name" binding:"required"`
	Email             string        `json:"email"`
	Phone             string        `json:"phone"`
	WhatsApp          string        `json:"whatsapp"`
	IdentificationNum string        `json:"identification_num"`
	Address           string        `json:"address"`
	Source            string        `json:"source"`
	Status            string        `json:"status"`
	ResponsibleID     string        `json:"responsible_id"`
	Tags              []string      `json:"tags"`
	Person            shared.Person `json:"person"`
	IsActive          bool          `json:"is_active"`
	Comment           string        `json:"comment"`
}

type UpdateClientRequest struct {
	Name              string        `json:"name"`
	Email             string        `json:"email"`
	Phone             string        `json:"phone"`
	WhatsApp          string        `json:"whatsapp"`
	IdentificationNum string        `json:"identification_num"`
	Address           string        `json:"address"`
	Source            string        `json:"source"`
	Status            string        `json:"status"`
	ResponsibleID     string        `json:"responsible_id"`
	Tags              []string      `json:"tags"`
	Person            shared.Person `json:"person"`
	IsActive          bool          `json:"is_active"`
	Comment           string        `json:"comment"`
}
