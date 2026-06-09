package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
	"github.com/gin-gonic/gin"
)

type CommunicationHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewCommunicationHandler(deps *http_deps.Dependencies, logs *slog.Logger) *CommunicationHandler {
	return &CommunicationHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/communications communications createCommunication
// Security:
//   Bearer:
// responses:
//   201: getCommunicationResponse
//   400: errorResponse
func (h *CommunicationHandler) Create(c *gin.Context) {
	var req http_requests.CreateCommunicationRequest
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

	comm := &communication.Communication{
		UserID:   userID,
		ClientID: req.ClientID,
	}

	res, err := h.deps.App.Commands.CreateCommunication.Handle(c.Request.Context(), command.CreateCommunicationCommand{
		Communication: comm,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/communications/{id} communications getCommunication
// Security:
//   Bearer:
// responses:
//   200: getCommunicationResponse
//   404: errorResponse
func (h *CommunicationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetCommunicationByID.Handle(c.Request.Context(), query.FetchCommunicationByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/communications/{id} communications updateCommunication
// Security:
//   Bearer:
// responses:
//   200: getCommunicationResponse
//   400: errorResponse
//   404: errorResponse
func (h *CommunicationHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateCommunicationRequest
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

	comm := &communication.Communication{
		ID:       id,
		UserID:   userID,
		ClientID: req.ClientID,
	}

	res, err := h.deps.App.Commands.UpdateCommunication.Handle(c.Request.Context(), command.UpdateCommunicationCommand{
		Communication: comm,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/communications/{id} communications deleteCommunication
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *CommunicationHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteCommunication.Handle(c.Request.Context(), command.DeleteCommunicationCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/communications communications listCommunications
// Security:
//   Bearer:
// responses:
//   200: listCommunicationsResponse
//   500: errorResponse
func (h *CommunicationHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListCommunications.Handle(c.Request.Context(), query.FetchCommunicationList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
