package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
)

type FetchInvoiceByID struct {
	ID string
}

type FetchInvoiceByIDHandler struct {
	repo invoice.InvoiceRepositoryInterface
}

func NewFetchInvoiceByIDHandler(repo invoice.InvoiceRepositoryInterface) *FetchInvoiceByIDHandler {
	return &FetchInvoiceByIDHandler{repo: repo}
}

func (h *FetchInvoiceByIDHandler) Handle(ctx context.Context, q FetchInvoiceByID) (*invoice.Invoice, error) {
	return h.repo.GetByID(q.ID)
}

type FetchInvoiceList struct{}

type FetchInvoiceListHandler struct {
	repo invoice.InvoiceRepositoryInterface
}

func NewFetchInvoiceListHandler(repo invoice.InvoiceRepositoryInterface) *FetchInvoiceListHandler {
	return &FetchInvoiceListHandler{repo: repo}
}

func (h *FetchInvoiceListHandler) Handle(ctx context.Context, q FetchInvoiceList) ([]*invoice.Invoice, error) {
	return h.repo.List()
}
