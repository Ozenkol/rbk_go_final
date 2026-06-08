package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewDocumentHandler(deps *http_deps.Dependencies, logs *slog.Logger) *DocumentHandler {
	return &DocumentHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/documents documents createDocument
// Create a new document.
// security:
//   - Bearer:
// responses:
//   201: getDocumentResponse
//   400: errorResponse
func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req http_requests.CreateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	userID, err := h.deps.App.Services.AuthService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	d := &document.Document{
		UserID:         userID,
		ClientID:       req.ClientID,
		CompanyID:      req.CompanyID,
		Type:           req.Type,
		Number:         req.Number,
		IssuedBy:       req.IssuedBy,
		IssuedDate:     req.IssuedDate,
		ExpirationDate: req.ExpirationDate,
	}

	res, err := h.deps.App.Commands.CreateDocument.Handle(c.Request.Context(), command.CreateDocumentCommand{
		Document: d,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/documents/{id} documents getDocument
// Get a document by ID.
// security:
//   - Bearer:
// responses:
//   200: getDocumentResponse
//   404: errorResponse
func (h *DocumentHandler) GetDocument(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetDocumentByID.Handle(c.Request.Context(), query.FetchDocumentByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/documents/{id} documents updateDocument
// Update a document by ID.
// security:
//   - Bearer:
// responses:
//   200: getDocumentResponse
//   400: errorResponse
func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	userID, err := h.deps.App.Services.AuthService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	d := &document.Document{
		ID:             id,
		UserID:         userID,
		ClientID:       req.ClientID,
		CompanyID:      req.CompanyID,
		Type:           req.Type,
		Number:         req.Number,
		IssuedBy:       req.IssuedBy,
		IssuedDate:     req.IssuedDate,
		ExpirationDate: req.ExpirationDate,
	}

	res, err := h.deps.App.Commands.UpdateDocument.Handle(c.Request.Context(), command.UpdateDocumentCommand{
		Document: d,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/documents/{id} documents deleteDocument
// Delete a document by ID.
// security:
//   - Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteDocument.Handle(c.Request.Context(), command.DeleteDocumentCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
