package communication

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Communication struct {
	ID        string                   `json:"id"`
	UserID    string                   `json:"user_id"`
	CompanyID string                   `json:"company_id"`
	ClientID  string                   `json:"client_id"`
	DealID    string                   `json:"deal_id"`
	Type      shared.CommunicationType `json:"type"`
	Content   string                   `json:"content"`
	CreatedAt int64                    `json:"created_at"`
}
