package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"gorm.io/gorm"
)

type InvoiceModel struct {
	ID   string `gorm:"primaryKey"`
}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) (invoice.InvoiceRepositoryInterface, error) {
	if err := db.AutoMigrate(&InvoiceModel{}); err != nil {
		return nil, err
	}
	return &InvoiceRepository{db: db}, nil
}

func (r *InvoiceRepository) Create(i *invoice.Invoice) (*invoice.Invoice, error) {
	model := &InvoiceModel{ID: i.ID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &invoice.Invoice{ID: model.ID}, nil
}

func (r *InvoiceRepository) GetByID(id string) (*invoice.Invoice, error) {
	var model InvoiceModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice.Invoice{ID: model.ID}, nil
}

func (r *InvoiceRepository) Update(i *invoice.Invoice) (*invoice.Invoice, error) {
	model := &InvoiceModel{ID: i.ID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &invoice.Invoice{ID: model.ID}, nil
}

func (r *InvoiceRepository) Delete(id string) error {
	return r.db.Delete(&InvoiceModel{}, "id = ?", id).Error
}

func (r *InvoiceRepository) List() ([]*invoice.Invoice, error) {
	var models []InvoiceModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	invoices := make([]*invoice.Invoice, len(models))
	for idx, m := range models {
		invoices[idx] = &invoice.Invoice{ID: m.ID}
	}
	return invoices, nil
}
