package handlers

import (
	"log/slog"

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

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Create document - protected route"})
	h.logs.Info("Create document - protected route")
}

func (h *DocumentHandler) GetDocument(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get document - protected route"})
	h.logs.Info("Get document - protected route")
}

func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update document - protected route"})
	h.logs.Info("Update document - protected route")
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete document - protected route"})
	h.logs.Info("Delete document - protected route")
}

