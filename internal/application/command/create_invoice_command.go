package command

import "github.com/Ozenkol/rbk-go-final/internal/domain/invoice"

type CreateInvoiceCommand struct {
	UserID   string
	ClientID string
	Amount   float64
}

type CreateInvoiceHandler struct {
	invoiceRepo invoice.InvoiceRepositoryInterface
}

func NewCreateInvoiceHandler(invoiceRepo invoice.InvoiceRepositoryInterface) *CreateInvoiceHandler {
	return &CreateInvoiceHandler{
		invoiceRepo: invoiceRepo,
	}
}

func (h *CreateInvoiceHandler) Handle(cmd CreateInvoiceCommand) (*invoice.Invoice, error) {
	newInvoice := &invoice.Invoice{
		ClientID:   cmd.ClientID,
		DocumentID: "",
	}
	h.invoiceRepo.Create(newInvoice)

	return newInvoice, nil
}