package contract

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Contract struct {
	ID             string                `json:"id"`
	UserID         string                `json:"user_id"`
	ClientID       string                `json:"client_id"`
	DealID         string                `json:"deal_id"`
	CompanyID      string                `json:"company_id"`
	Number         string                `json:"number"`
	Status         shared.ContractStatus `json:"status"`
	ValidFrom      int64                 `json:"valid_from"`
	ValidUntil     int64                 `json:"valid_until"`
	FileID         string                `json:"file_id"`
	CreatedAt      int64                 `json:"created_at"`
	UpdatedAt      int64                 `json:"updated_at"`
}
