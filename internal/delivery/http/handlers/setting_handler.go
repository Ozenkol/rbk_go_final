package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type SettingHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewSettingHandler(deps *http_deps.Dependencies, logs *slog.Logger) *SettingHandler {
	return &SettingHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/settings settings createSetting
func (h *SettingHandler) Create(c *gin.Context) {
	var s setting.Setting
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateSetting.Handle(c.Request.Context(), command.CreateSettingCommand{Setting: &s})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/settings/{id} settings getSetting
func (h *SettingHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetSettingByID.Handle(c.Request.Context(), query.FetchSettingByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/settings/{id} settings updateSetting
func (h *SettingHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var s setting.Setting
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.ID = id
	res, err := h.deps.App.Commands.UpdateSetting.Handle(c.Request.Context(), command.UpdateSettingCommand{Setting: &s})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/settings/{id} settings deleteSetting
func (h *SettingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteSetting.Handle(c.Request.Context(), command.DeleteSettingCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/settings settings listSettings
func (h *SettingHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListSettings.Handle(c.Request.Context(), query.FetchSettingList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
