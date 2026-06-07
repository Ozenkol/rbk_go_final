package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type MeetingHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewMeetingHandler(deps *http_deps.Dependencies, logs *slog.Logger) *MeetingHandler {
	return &MeetingHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/meetings meetings createMeeting
func (h *MeetingHandler) Create(c *gin.Context) {
	var m meeting.Meeting
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateMeeting.Handle(c.Request.Context(), command.CreateMeetingCommand{Meeting: &m})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/meetings/{id} meetings getMeeting
func (h *MeetingHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetMeetingByID.Handle(c.Request.Context(), query.FetchMeetingByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/meetings/{id} meetings updateMeeting
func (h *MeetingHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var m meeting.Meeting
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	m.ID = id
	res, err := h.deps.App.Commands.UpdateMeeting.Handle(c.Request.Context(), command.UpdateMeetingCommand{Meeting: &m})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/meetings/{id} meetings deleteMeeting
func (h *MeetingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteMeeting.Handle(c.Request.Context(), command.DeleteMeetingCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/meetings meetings listMeetings
func (h *MeetingHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListMeetings.Handle(c.Request.Context(), query.FetchMeetingList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
