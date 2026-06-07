package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"gorm.io/gorm"
)

type InvoiceItemModel struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	InvoiceID string
	ProductID string
	Amount  int
	Description     string
}

type InvoiceModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	ClientID     string
	DocumentID   string
	Items []InvoiceItemModel `gorm:"foreignKey:InvoiceID"`
}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) (invoice.InvoiceRepositoryInterface, error) {
	if err := db.AutoMigrate(&InvoiceModel{}, &InvoiceItemModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &InvoiceRepository{db: db}, nil
}

func (r *InvoiceRepository) Create(invoice *invoice.Invoice) error {	
	invoiceModel := toInvoiceModel(invoice)
	return r.db.Create(invoiceModel).Error
}

func (r *InvoiceRepository) GetByID(id string) (*invoice.Invoice, error) {
	var model InvoiceModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}

	invoice := toInvoiceDomain(&model)

	var items []InvoiceItemModel
	if err := r.db.Where("invoice_id = ?", id).Find(&items).Error; err != nil {
		return nil, err
	}
	for i, itemModel := range items {
		invoice.InvoiceItems[i] = *toInvoiceItemDomain(&itemModel)
	}
	return invoice, nil
}

func (r *InvoiceRepository) Update(invoice *invoice.Invoice) (*invoice.Invoice, error) {
	invoiceModel := toInvoiceModel(invoice)
	err := r.db.Save(invoiceModel).Error
	if err != nil {
		return nil, err
	}
	return toInvoiceDomain(invoiceModel), nil
}

func (r *InvoiceRepository) Delete(id string) error {	
	return r.db.Where("id = ?", id).Delete(&InvoiceModel{}).Error
}

func toInvoiceModel(invoice *invoice.Invoice) *InvoiceModel {
	return &InvoiceModel{
		ID:         invoice.ID,
		ClientID:   invoice.ClientID,
		DocumentID: invoice.DocumentID,
	}
}

func toInvoiceDomain(invoiceModel *InvoiceModel) *invoice.Invoice {
	return &invoice.Invoice{
		ID:         invoiceModel.ID,
		ClientID:   invoiceModel.ClientID,
		DocumentID: invoiceModel.DocumentID,
	}
}

func toInvoiceItemModel(item *invoice.InvoiceItem) *InvoiceItemModel {
	return &InvoiceItemModel{
		ProductID: item.ProductID,
		Amount:  item.Amount,
		Description: item.Description,
	}
}

func toInvoiceItemDomain(itemModel *InvoiceItemModel) *invoice.InvoiceItem {
	return &invoice.InvoiceItem{
		InvoiceID: itemModel.InvoiceID,
		ProductID: itemModel.ProductID,
		Amount:  itemModel.Amount,
		Description: itemModel.Description,
	}
}	