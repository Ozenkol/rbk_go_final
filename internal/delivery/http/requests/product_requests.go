package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

// CreateProductRequest represents the request body for creating a product
// swagger:model CreateProductRequest
type CreateProductRequest struct {
	// Product name
	// example: Premium Subscription
	Name        string                  `json:"name"`
	// Product description
	// example: Access to all premium features
	Description string                  `json:"description"`
	// Thumbnail storage reference
	Thumbnail   shared.StorageReference `json:"thumbnail"`
	// Product price
	Price       shared.Money            `json:"price"`
}

// UpdateProductRequest represents the request body for updating a product
// swagger:model UpdateProductRequest
type UpdateProductRequest struct {
	// Product name
	// example: Premium Subscription
	Name        string                  `json:"name"`
	// Product description
	// example: Access to all premium features
	Description string                  `json:"description"`
	// Thumbnail storage reference
	Thumbnail   shared.StorageReference `json:"thumbnail"`
	// Product price
	Price       shared.Money            `json:"price"`
}
