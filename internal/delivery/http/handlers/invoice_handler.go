package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InvoiceHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewInvoiceHandler(deps *http_deps.Dependencies, logs *slog.Logger) *InvoiceHandler {
	return &InvoiceHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/invoices invoices createInvoice
// Create a new invoice.
// Security:
//   Bearer:
// responses:
//   201: body:Invoice
//   400: body:errorResponse
func (h *InvoiceHandler) Create(c *gin.Context) {
	var req http_requests.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	userID, companyID, err := h.deps.App.Services.AuthService.GetAuthInfoFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	inv := &invoice.Invoice{
		ID:        uuid.New().String(),
		UserID:    userID,
		ClientID:  req.ClientID,
		DealID:    req.DealID,
		CompanyID: companyID,
		Number:    req.Number,
		Status:    shared.InvoiceStatusDraft,
		DueDate:   req.DueDate,
	}

	res, err := h.deps.App.Commands.CreateInvoice.Handle(c.Request.Context(), command.CreateInvoiceCommand{Invoice: inv})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/invoices/{id} invoices getInvoice
// Get invoice by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Invoice
//   404: body:errorResponse
func (h *InvoiceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetInvoiceByID.Handle(c.Request.Context(), query.FetchInvoiceByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/invoices invoices listInvoices
// List all invoices.
// Security:
//   Bearer:
// responses:
//   200: body:[]Invoice
func (h *InvoiceHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListInvoices.Handle(c.Request.Context(), query.FetchInvoiceList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/invoices/{id} invoices updateInvoice
// Update an existing invoice.
// Security:
//   Bearer:
// responses:
//   200: body:Invoice
//   400: body:errorResponse
//   404: body:errorResponse
func (h *InvoiceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inv, err := h.deps.App.Queries.GetInvoiceByID.Handle(c.Request.Context(), query.FetchInvoiceByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	if req.Status != "" { inv.Status = shared.InvoiceStatus(req.Status) }
	if req.PaidAmount != 0 { inv.PaidAmount = req.PaidAmount }

	res, err := h.deps.App.Commands.UpdateInvoice.Handle(c.Request.Context(), command.UpdateInvoiceCommand{Invoice: inv})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/invoices/{id} invoices deleteInvoice
// Delete invoice by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *InvoiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteInvoice.Handle(c.Request.Context(), command.DeleteInvoiceCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
