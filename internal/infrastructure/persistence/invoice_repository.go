package persistence

import (
	"encoding/json"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type InvoiceModel struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	UserID      string
	ClientID    string
	DealID      string
	CompanyID   string
	Number      string
	Status      string
	TotalAmount float64
	PaidAmount  float64
	DueDate     int64
	ItemsJSON   string `gorm:"type:text"`
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
	model, err := toInvoiceModel(i)
	if err != nil {
		return nil, err
	}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toInvoiceDomain(model)
}

func (r *InvoiceRepository) GetByID(id string) (*invoice.Invoice, error) {
	var model InvoiceModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toInvoiceDomain(&model)
}

func (r *InvoiceRepository) Update(i *invoice.Invoice) (*invoice.Invoice, error) {
	model, err := toInvoiceModel(i)
	if err != nil {
		return nil, err
	}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toInvoiceDomain(model)
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
	for i, m := range models {
		inv, err := toInvoiceDomain(&m)
		if err != nil {
			return nil, err
		}
		invoices[i] = inv
	}
	return invoices, nil
}

func toInvoiceModel(i *invoice.Invoice) (*InvoiceModel, error) {
	itemsJSON, err := json.Marshal(i.InvoiceItems)
	if err != nil {
		return nil, err
	}
	return &InvoiceModel{
		ID:          i.ID,
		UserID:      i.UserID,
		ClientID:    i.ClientID,
		DealID:      i.DealID,
		CompanyID:   i.CompanyID,
		Number:      i.Number,
		Status:      string(i.Status),
		TotalAmount: i.TotalAmount,
		PaidAmount:  i.PaidAmount,
		DueDate:     i.DueDate,
		ItemsJSON:   string(itemsJSON),
	}, nil
}

func toInvoiceDomain(m *InvoiceModel) (*invoice.Invoice, error) {
	var items []invoice.InvoiceItem
	if m.ItemsJSON != "" {
		if err := json.Unmarshal([]byte(m.ItemsJSON), &items); err != nil {
			return nil, err
		}
	}
	return &invoice.Invoice{
		ID:           m.ID,
		UserID:       m.UserID,
		ClientID:     m.ClientID,
		DealID:       m.DealID,
		CompanyID:    m.CompanyID,
		Number:       m.Number,
		Status:       shared.InvoiceStatus(m.Status),
		TotalAmount:  m.TotalAmount,
		PaidAmount:   m.PaidAmount,
		DueDate:      m.DueDate,
		CreatedAt:    m.CreatedAt.Unix(),
		UpdatedAt:    m.UpdatedAt.Unix(),
		InvoiceItems: items,
	}, nil
}
