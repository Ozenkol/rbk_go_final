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

// swagger:route POST /api/v1/offers offers createOffer
//
// Create a new offer.
//
// Consumes:
// - application/json
// Produces:
// - application/json
func (h *OfferHandler) CreateOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create offer - protected route"})
}

// swagger:route GET /api/v1/offers/{id} offers getOffer
//
// Get an offer by ID.
//
// Produces:
// - application/json
func (h *OfferHandler) GetOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get offer - protected route"})
}

// swagger:route PUT /api/v1/offers/{id} offers updateOffer
//
// Update an offer by ID.
//
// Consumes:
// - application/json
// Produces:
// - application/json
func (h *OfferHandler) UpdateOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update offer - protected route"})
}

// swagger:route DELETE /api/v1/offers/{id} offers deleteOffer
//
// Delete an offer by ID.
//
// Produces:
// - application/json
func (h *OfferHandler) DeleteOffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete offer - protected route"})
}

