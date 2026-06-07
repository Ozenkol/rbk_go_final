package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
)

type CreateInvoiceCommand struct {
	Invoice *invoice.Invoice
}

type CreateInvoiceHandler struct {
	repo invoice.InvoiceRepositoryInterface
}

func NewCreateInvoiceHandler(repo invoice.InvoiceRepositoryInterface) *CreateInvoiceHandler {
	return &CreateInvoiceHandler{repo: repo}
}

func (h *CreateInvoiceHandler) Handle(ctx context.Context, cmd CreateInvoiceCommand) (*invoice.Invoice, error) {
	return h.repo.Create(cmd.Invoice)
}

type UpdateInvoiceCommand struct {
	Invoice *invoice.Invoice
}

type UpdateInvoiceHandler struct {
	repo invoice.InvoiceRepositoryInterface
}

func NewUpdateInvoiceHandler(repo invoice.InvoiceRepositoryInterface) *UpdateInvoiceHandler {
	return &UpdateInvoiceHandler{repo: repo}
}

func (h *UpdateInvoiceHandler) Handle(ctx context.Context, cmd UpdateInvoiceCommand) (*invoice.Invoice, error) {
	return h.repo.Update(cmd.Invoice)
}

type DeleteInvoiceCommand struct {
	ID string
}

type DeleteInvoiceHandler struct {
	repo invoice.InvoiceRepositoryInterface
}

func NewDeleteInvoiceHandler(repo invoice.InvoiceRepositoryInterface) *DeleteInvoiceHandler {
	return &DeleteInvoiceHandler{repo: repo}
}

func (h *DeleteInvoiceHandler) Handle(ctx context.Context, cmd DeleteInvoiceCommand) error {
	return h.repo.Delete(cmd.ID)
}
