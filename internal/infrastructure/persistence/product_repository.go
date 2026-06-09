package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	ID             string `gorm:"primaryKey"`
	UserID         string
	CompanyID      string
	Name           string
	Description    string
	Type           string
	Category       string
	SKU            string
	Price          float64
	Currency       string
	Unit           string
	IsActive       bool
	StorageService string
	StorageURL     string
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) (product.ProductRepositoryInterface, error) {
	if err := db.AutoMigrate(&ProductModel{}); err != nil {
		return nil, err
	}
	return &ProductRepository{db: db}, nil
}

func (r *ProductRepository) Create(p *product.Product) (*product.Product, error) {
	model := toProductModel(p)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toProductDomain(model), nil
}

func (r *ProductRepository) GetByID(id string) (*product.Product, error) {
	var model ProductModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toProductDomain(&model), nil
}

func (r *ProductRepository) Update(p *product.Product) (*product.Product, error) {
	model := toProductModel(p)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toProductDomain(model), nil
}

func (r *ProductRepository) Delete(id string) error {
	return r.db.Delete(&ProductModel{}, "id = ?", id).Error
}

func (r *ProductRepository) List() ([]*product.Product, error) {
	var models []ProductModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	products := make([]*product.Product, len(models))
	for i, m := range models {
		products[i] = toProductDomain(&m)
	}
	return products, nil
}

func toProductModel(p *product.Product) *ProductModel {
	return &ProductModel{
		ID:             p.ID,
		UserID:         p.UserID,
		CompanyID:      p.CompanyID,
		Name:           p.Name,
		Description:    p.Description,
		Type:           p.Type,
		Category:       p.Category,
		SKU:            p.SKU,
		Price:          p.Price,
		Currency:       p.Currency,
		Unit:           p.Unit,
		IsActive:       p.IsActive,
		StorageService: p.Thumbnail.ServiceName,
		StorageURL:     p.Thumbnail.URL,
	}
}

func toProductDomain(m *ProductModel) *product.Product {
	return &product.Product{
		ID:          m.ID,
		UserID:      m.UserID,
		CompanyID:   m.CompanyID,
		Name:        m.Name,
		Description: m.Description,
		Type:        m.Type,
		Category:    m.Category,
		SKU:         m.SKU,
		Price:       m.Price,
		Currency:    m.Currency,
		Unit:        m.Unit,
		IsActive:    m.IsActive,
		Thumbnail:   shared.StorageReference{ServiceName: m.StorageService, URL: m.StorageURL},
	}
}
