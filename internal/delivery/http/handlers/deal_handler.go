package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/deal"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DealHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewDealHandler(deps *http_deps.Dependencies, logs *slog.Logger) *DealHandler {
	return &DealHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/deals deals createDeal
// Create a new deal.
// Security:
//   Bearer:
// responses:
//   201: body:Deal
//   400: body:errorResponse
func (h *DealHandler) CreateDeal(c *gin.Context) {
	var req http_requests.CreateDealRequest
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

	d := &deal.Deal{
		UserID:        userID,
		CompanyID:     companyID,
		ID:            uuid.New().String(),
		ClientID:      req.ClientID,
		ResponsibleID: req.ResponsibleID,
		Title:         req.Title,
		Stage:         shared.DealStage(req.Stage),
		Amount:        req.Amount,
		Currency:      req.Currency,
		Probability:   req.Probability,
		Deadline:      req.Deadline,
	}

	res, err := h.deps.App.Commands.CreateDeal.Handle(c.Request.Context(), command.CreateDealCommand{Deal: d})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/deals/{id} deals getDeal
// Get deal by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Deal
//   404: body:errorResponse
func (h *DealHandler) GetDeal(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetDealByID.Handle(c.Request.Context(), query.FetchDealByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Deal not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/deals deals listDeals
// List all deals.
// Security:
//   Bearer:
// responses:
//   200: body:[]Deal
func (h *DealHandler) ListDeals(c *gin.Context) {
	res, err := h.deps.App.Queries.ListDeals.Handle(c.Request.Context(), query.FetchDealList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/deals/{id} deals updateDeal
// Update an existing deal.
// Security:
//   Bearer:
// responses:
//   200: body:Deal
//   400: body:errorResponse
//   404: body:errorResponse
func (h *DealHandler) UpdateDeal(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateDealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := h.deps.App.Queries.GetDealByID.Handle(c.Request.Context(), query.FetchDealByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Deal not found"})
		return
	}

	if req.Title != "" { d.Title = req.Title }
	if req.Stage != "" { d.Stage = shared.DealStage(req.Stage) }
	if req.Amount != 0 { d.Amount = req.Amount }
	if req.Currency != "" { d.Currency = req.Currency }
	if req.Probability != 0 { d.Probability = req.Probability }
	if req.Deadline != 0 { d.Deadline = req.Deadline }
	if req.ResponsibleID != "" { d.ResponsibleID = req.ResponsibleID }

	res, err := h.deps.App.Commands.UpdateDeal.Handle(c.Request.Context(), command.UpdateDealCommand{Deal: d})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/deals/{id} deals deleteDeal
// Delete deal by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *DealHandler) DeleteDeal(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteDeal.Handle(c.Request.Context(), command.DeleteDealCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
