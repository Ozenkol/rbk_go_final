package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
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
	var d document.Document
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateDocument.Handle(c.Request.Context(), command.CreateDocumentCommand{Document: &d})
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
	var d document.Document
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d.ID = id
	res, err := h.deps.App.Commands.UpdateDocument.Handle(c.Request.Context(), command.UpdateDocumentCommand{Document: &d})
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
