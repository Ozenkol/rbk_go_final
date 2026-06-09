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
	"github.com/google/uuid"
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
// Security:
//   Bearer:
// responses:
//   201: body:Document
//   400: body:errorResponse
func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req http_requests.CreateDocumentRequest
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

	doc := &document.Document{
		ID:             uuid.New().String(),
		UserID:         userID,
		ClientID:       req.ClientID,
		DealID:         req.DealID,
		ContractID:     req.ContractID,
		CompanyID:      companyID,
		Type:           req.Type,
		Number:         req.Number,
		IssuedBy:       req.IssuedBy,
		IssuedDate:     req.IssuedDate,
		ExpirationDate: req.ExpirationDate,
	}

	res, err := h.deps.App.Commands.CreateDocument.Handle(c.Request.Context(), command.CreateDocumentCommand{Document: doc})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/documents/{id} documents getDocument
// Get document by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Document
//   404: body:errorResponse
func (h *DocumentHandler) GetDocument(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetDocumentByID.Handle(c.Request.Context(), query.FetchDocumentByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/documents documents listDocuments
// List all documents.
// Security:
//   Bearer:
// responses:
//   200: body:[]DocumentWithURL
func (h *DocumentHandler) ListDocuments(c *gin.Context) {
	res, err := h.deps.App.Queries.ListDocuments.Handle(c.Request.Context(), query.FetchDocumentList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/documents/{id} documents updateDocument
// Update an existing document.
// Security:
//   Bearer:
// responses:
//   200: body:Document
//   400: body:errorResponse
//   404: body:errorResponse
func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resByID, err := h.deps.App.Queries.GetDocumentByID.Handle(c.Request.Context(), query.FetchDocumentByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}
	doc := resByID.Document

	if req.Type != "" { doc.Type = req.Type }
	if req.Number != "" { doc.Number = req.Number }
	if req.IssuedBy != "" { doc.IssuedBy = req.IssuedBy }
	if req.IssuedDate != 0 { doc.IssuedDate = req.IssuedDate }
	if req.ExpirationDate != 0 { doc.ExpirationDate = req.ExpirationDate }

	res, err := h.deps.App.Commands.UpdateDocument.Handle(c.Request.Context(), command.UpdateDocumentCommand{Document: doc})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/documents/{id} documents deleteDocument
// Delete document by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteDocument.Handle(c.Request.Context(), command.DeleteDocumentCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
