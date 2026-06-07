package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewInvoiceHandler(deps *http_deps.Dependencies, logs *slog.Logger) *InvoiceHandler {
	return &InvoiceHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/invoices invoices createInvoice
func (h *InvoiceHandler) Create(c *gin.Context) {
	var i invoice.Invoice
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateInvoice.Handle(c.Request.Context(), command.CreateInvoiceCommand{Invoice: &i})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/invoices/{id} invoices getInvoice
func (h *InvoiceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetInvoiceByID.Handle(c.Request.Context(), query.FetchInvoiceByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/invoices/{id} invoices updateInvoice
func (h *InvoiceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var i invoice.Invoice
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i.ID = id
	res, err := h.deps.App.Commands.UpdateInvoice.Handle(c.Request.Context(), command.UpdateInvoiceCommand{Invoice: &i})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/invoices/{id} invoices deleteInvoice
func (h *InvoiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteInvoice.Handle(c.Request.Context(), command.DeleteInvoiceCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/invoices invoices listInvoices
func (h *InvoiceHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListInvoices.Handle(c.Request.Context(), query.FetchInvoiceList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
