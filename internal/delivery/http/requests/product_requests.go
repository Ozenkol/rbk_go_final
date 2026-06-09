package http_requests

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Unit        string  `json:"unit"`
	IsActive    bool    `json:"is_active"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsActive    bool    `json:"is_active"`
}
