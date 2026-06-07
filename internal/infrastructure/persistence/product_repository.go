package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"gorm.io/gorm"
)

type ProductModel struct {
	ID   string `gorm:"primaryKey"`
	Name string
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
	model := &ProductModel{ID: p.ID, Name: p.Name}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &product.Product{ID: model.ID, Name: model.Name}, nil
}

func (r *ProductRepository) GetByID(id string) (*product.Product, error) {
	var model ProductModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product.Product{ID: model.ID, Name: model.Name}, nil
}

func (r *ProductRepository) Update(p *product.Product) (*product.Product, error) {
	model := &ProductModel{ID: p.ID, Name: p.Name}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &product.Product{ID: model.ID, Name: model.Name}, nil
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
		products[i] = &product.Product{ID: m.ID, Name: m.Name}
	}
	return products, nil
}
