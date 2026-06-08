package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/invoice"
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
// Создать новый счет.
// Security:
//   Bearer:
// responses:
//   201: getInvoiceResponse
//   400: errorResponse
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

	i := &invoice.Invoice{
		UserID:     userID,
		CompanyID:  companyID,
		ClientID:   req.ClientID,
		DocumentID: req.DocumentID,
	}

	res, err := h.deps.App.Commands.CreateInvoice.Handle(c.Request.Context(), command.CreateInvoiceCommand{
		Invoice: i,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/invoices/{id} invoices getInvoice
// Получить счет по ID.
// Security:
//   Bearer:
// responses:
//   200: getInvoiceResponse
//   404: errorResponse
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
// Список всех счетов.
// Security:
//   Bearer:
// responses:
//   200: []getInvoiceResponse
//   500: errorResponse
func (h *InvoiceHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListInvoices.Handle(c.Request.Context(), query.FetchInvoiceList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/invoices/{id} invoices updateInvoice
// Обновить счет по ID.
// Security:
//   Bearer:
// responses:
//   200: getInvoiceResponse
//   400: errorResponse
func (h *InvoiceHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateInvoiceRequest
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

	i := &invoice.Invoice{
		ID:         id,
		UserID:     userID,
		CompanyID:  companyID,
		ClientID:   req.ClientID,
		DocumentID: req.DocumentID,
	}

	res, err := h.deps.App.Commands.UpdateInvoice.Handle(c.Request.Context(), command.UpdateInvoiceCommand{
		Invoice: i,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/invoices/{id} invoices deleteInvoice
// Удалить счет по ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *InvoiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteInvoice.Handle(c.Request.Context(), command.DeleteInvoiceCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
