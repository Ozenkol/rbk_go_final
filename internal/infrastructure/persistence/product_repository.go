package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type ProductModel struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Thumbnail   shared.StorageReference
	Price       shared.Money
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) (product.ProductRepositoryInterface, error) {
	if err := db.AutoMigrate(&ProductModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &ProductRepository{db: db}, nil
}

func (r *ProductRepository) GetProductByID(id string) (*product.Product, error) {
	var productModel ProductModel
	if err := r.db.First(&productModel, id).Error; err != nil {
		return nil, err
	}
	return fromProductModel(&productModel), nil
}

func toProductModel(p *product.Product) *ProductModel {
	return &ProductModel{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func fromProductModel(productModel *ProductModel) *product.Product {
	return &product.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Price:       productModel.Price,
	}
}
