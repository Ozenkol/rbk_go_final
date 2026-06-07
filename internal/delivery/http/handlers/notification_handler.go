package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewNotificationHandler(deps *http_deps.Dependencies, logs *slog.Logger) *NotificationHandler {
	return &NotificationHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/notifications notifications createNotification
func (h *NotificationHandler) Create(c *gin.Context) {
	var n notification.Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateNotification.Handle(c.Request.Context(), command.CreateNotificationCommand{Notification: &n})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/notifications/{id} notifications getNotification
func (h *NotificationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetNotificationByID.Handle(c.Request.Context(), query.FetchNotificationByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/notifications/{id} notifications updateNotification
func (h *NotificationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var n notification.Notification
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	n.ID = id
	res, err := h.deps.App.Commands.UpdateNotification.Handle(c.Request.Context(), command.UpdateNotificationCommand{Notification: &n})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/notifications/{id} notifications deleteNotification
func (h *NotificationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteNotification.Handle(c.Request.Context(), command.DeleteNotificationCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/notifications notifications listNotifications
func (h *NotificationHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListNotifications.Handle(c.Request.Context(), query.FetchNotificationList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
