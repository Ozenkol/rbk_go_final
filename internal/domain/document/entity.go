package document

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Document struct {
	ID               string                  `json:"id"`
	UserID           string                  `json:"user_id"`
	ClientID         string                  `json:"client_id"`
	DealID           string                  `json:"deal_id"`
	ContractID       string                  `json:"contract_id"`
	CompanyID        string                  `json:"company_id"`
	Type             string                  `json:"type"` // Identity, Registration, etc.
	Number           string                  `json:"number"`
	IssuedBy         string                  `json:"issued_by"`
	StorageReference shared.StorageReference `json:"storage_reference"`
	IssuedDate       int64                   `json:"issued_date"`
	ExpirationDate   int64                   `json:"expiration_date"`
	CreatedAt        int64                   `json:"created_at"`
}
