package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
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
// Security:
//   Bearer:
// responses:
//   201: getNotificationResponse
//   400: errorResponse
func (h *NotificationHandler) Create(c *gin.Context) {
	var req http_requests.CreateNotificationRequest
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

	n := &notification.Notification{
		UserID:   userID,
		ClientID: req.ClientID,
		Message:  req.Message,
	}

	res, err := h.deps.App.Commands.CreateNotification.Handle(c.Request.Context(), command.CreateNotificationCommand{Notification: n})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/notifications/{id} notifications getNotification
// Security:
//   Bearer:
// responses:
//   200: getNotificationResponse
//   404: errorResponse
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
// Security:
//   Bearer:
// responses:
//   200: getNotificationResponse
//   400: errorResponse
//   404: errorResponse
func (h *NotificationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateNotificationRequest
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

	n := &notification.Notification{
		ID:       id,
		UserID:   userID,
		ClientID: req.ClientID,
		Message:  req.Message,
	}

	res, err := h.deps.App.Commands.UpdateNotification.Handle(c.Request.Context(), command.UpdateNotificationCommand{Notification: n})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/notifications/{id} notifications deleteNotification
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
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
// Security:
//   Bearer:
// responses:
//   200: listNotificationsResponse
//   500: errorResponse
func (h *NotificationHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListNotifications.Handle(c.Request.Context(), query.FetchNotificationList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
