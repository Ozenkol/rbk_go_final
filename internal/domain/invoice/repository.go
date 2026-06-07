package invoice

type InvoiceRepositoryInterface interface {
	Create(invoice *Invoice) error
	GetByID(id string) (*Invoice, error)
	Update(invoice *Invoice) (*Invoice, error)
	Delete(id string) error
}