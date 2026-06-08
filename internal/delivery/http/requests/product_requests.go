package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type CreateProductRequest struct {
	Name        string                  `json:"name"`
	CompanyID   string                  `json:"company_id"`
	Description string                  `json:"description"`
	Thumbnail   shared.StorageReference `json:"thumbnail"`
	Price       shared.Money            `json:"price"`
}

type UpdateProductRequest struct {
	Name        string                  `json:"name"`
	CompanyID   string                  `json:"company_id"`
	Description string                  `json:"description"`
	Thumbnail   shared.StorageReference `json:"thumbnail"`
	Price       shared.Money            `json:"price"`
}
