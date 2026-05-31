package product

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Product struct {
	ID          string
	Name        string
	Description string
	Thumbnail   shared.StorageReference
	Price       shared.Money
}