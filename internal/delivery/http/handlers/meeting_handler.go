package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
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
// Создать новую встречу.
// Security:
//   Bearer:
// responses:
//   201: getMeetingResponse
//   400: errorResponse
func (h *MeetingHandler) Create(c *gin.Context) {
	var req http_requests.CreateMeetingRequest
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

	m := &meeting.Meeting{
		UserID:          userID,
		CompanyID:       companyID,
		ClientID:        req.ClientID,
		Topic:           req.Topic,
		TimeSlot:        req.TimeSlot,
		MeetingProvider: req.MeetingProvider,
		Attendees:       req.Attendees,
	}

	res, err := h.deps.App.Commands.CreateMeeting.Handle(c.Request.Context(), command.CreateMeetingCommand{
		Meeting: m,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/meetings/{id} meetings getMeeting
// Получить встречу по ID.
// Security:
//   Bearer:
// responses:
//   200: getMeetingResponse
//   404: errorResponse
func (h *MeetingHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetMeetingByID.Handle(c.Request.Context(), query.FetchMeetingByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/meetings meetings listMeetings
// Список всех встреч.
// Security:
//   Bearer:
// responses:
//   200: []getMeetingResponse
//   500: errorResponse
func (h *MeetingHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListMeetings.Handle(c.Request.Context(), query.FetchMeetingList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/meetings/{id} meetings updateMeeting
// Обновить встречу по ID.
// Security:
//   Bearer:
// responses:
//   200: getMeetingResponse
//   400: errorResponse
func (h *MeetingHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateMeetingRequest
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

	m := &meeting.Meeting{
		ID:              id,
		UserID:          userID,
		CompanyID:       companyID,
		ClientID:        req.ClientID,
		Topic:           req.Topic,
		TimeSlot:        req.TimeSlot,
		MeetingProvider: req.MeetingProvider,
		Attendees:       req.Attendees,
	}

	res, err := h.deps.App.Commands.UpdateMeeting.Handle(c.Request.Context(), command.UpdateMeetingCommand{
		Meeting: m,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/meetings/{id} meetings deleteMeeting
// Удалить встречу по ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *MeetingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteMeeting.Handle(c.Request.Context(), command.DeleteMeetingCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
