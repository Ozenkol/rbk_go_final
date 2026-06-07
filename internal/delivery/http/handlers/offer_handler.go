package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
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
func (h *OfferHandler) CreateOffer(c *gin.Context) {
	var o offer.Offer
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateOffer.Handle(c.Request.Context(), command.CreateOfferCommand{Offer: &o})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/offers/{id} offers getOffer
func (h *OfferHandler) GetOffer(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetOfferByID.Handle(c.Request.Context(), query.FetchOfferByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/offers/{id} offers updateOffer
func (h *OfferHandler) UpdateOffer(c *gin.Context) {
	id := c.Param("id")
	var o offer.Offer
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	o.ID = id
	res, err := h.deps.App.Commands.UpdateOffer.Handle(c.Request.Context(), command.UpdateOfferCommand{Offer: &o})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/offers/{id} offers deleteOffer
func (h *OfferHandler) DeleteOffer(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteOffer.Handle(c.Request.Context(), command.DeleteOfferCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
