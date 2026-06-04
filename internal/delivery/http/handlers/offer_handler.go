package handlers

import (
	"log/slog"
	"net/http"

	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewOfferHandler(deps *http_deps.Dependencies, logs *slog.Logger) *OfferHandler {
	return &OfferHandler{deps: deps, logs: logs}
}

func (h *OfferHandler) CreateOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create offer - protected route"})
}

func (h *OfferHandler) GetOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get offer - protected route"})
}

func (h *OfferHandler) UpdateOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update offer - protected route"})
}

func (h *OfferHandler) DeleteOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete offer - protected route"})
}

