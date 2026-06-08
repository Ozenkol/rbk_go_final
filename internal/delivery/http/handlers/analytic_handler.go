package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
	"github.com/gin-gonic/gin"
)

type AnalyticHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewAnalyticHandler(deps *http_deps.Dependencies, logs *slog.Logger) *AnalyticHandler {
	return &AnalyticHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/analytics analytics createAnalytic
func (h *AnalyticHandler) Create(c *gin.Context) {
	var req http_requests.CreateAnalyticRequest
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

	a := &analytic.Analytic{
		UserID:   userID,
		ClientID: req.ClientID,
		Data:     req.Data,
	}

	res, err := h.deps.App.Commands.CreateAnalytic.Handle(c.Request.Context(), command.CreateAnalyticCommand{
		Analytic: a,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/analytics/{id} analytics getAnalytic
func (h *AnalyticHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetAnalyticByID.Handle(c.Request.Context(), query.FetchAnalyticByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/analytics/{id} analytics updateAnalytic
func (h *AnalyticHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateAnalyticRequest
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

	a := &analytic.Analytic{
		ID:       id,
		UserID:   userID,
		ClientID: req.ClientID,
		Data:     req.Data,
	}

	res, err := h.deps.App.Commands.UpdateAnalytic.Handle(c.Request.Context(), command.UpdateAnalyticCommand{
		Analytic: a,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/analytics/{id} analytics deleteAnalytic
func (h *AnalyticHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteAnalytic.Handle(c.Request.Context(), command.DeleteAnalyticCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/analytics analytics listAnalytics
func (h *AnalyticHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListAnalytics.Handle(c.Request.Context(), query.FetchAnalyticList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
