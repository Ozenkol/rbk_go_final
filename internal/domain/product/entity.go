package product

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Product struct {
	ID          string                  `json:"id"`
	UserID      string                  `json:"user_id"`
	CompanyID   string                  `json:"company_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Type        string                  `json:"type"` // Product / Service
	Category    string                  `json:"category"`
	SKU         string                  `json:"sku"`
	Price       float64                 `json:"price"`
	Currency    string                  `json:"currency"`
	Unit        string                  `json:"unit"`
	IsActive    bool                    `json:"is_active"`
	Thumbnail   shared.StorageReference `json:"thumbnail"`
}
