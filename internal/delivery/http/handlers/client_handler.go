package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewClientHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ClientHandler {
	return &ClientHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/clients clients createClient
func (h *ClientHandler) CreateClient(c *gin.Context) {
	var req http_requests.CreateClientRequest
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

	cl := &client.Client{
		UserID:    userID,
		CompanyID: req.CompanyID,
		Person:    req.Person,
		IsActive:  req.IsActive,
	}

	res, err := h.deps.App.Commands.CreateClient.Handle(c.Request.Context(), command.CreateClientCommand{
		Client: cl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/clients/{id} clients getClient
func (h *ClientHandler) GetClient(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetClientByID.Handle(c.Request.Context(), query.FetchClientByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/clients/{id} clients updateClient
func (h *ClientHandler) UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateClientRequest
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

	cl := &client.Client{
		ID:        id,
		UserID:    userID,
		CompanyID: req.CompanyID,
		Person:    req.Person,
		IsActive:  req.IsActive,
	}

	res, err := h.deps.App.Commands.UpdateClient.Handle(c.Request.Context(), command.UpdateClientCommand{
		Client: cl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/clients/{id} clients deleteClient
func (h *ClientHandler) DeleteClient(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteClient.Handle(c.Request.Context(), command.DeleteClientCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/clients clients listClients
func (h *ClientHandler) ListClients(c *gin.Context) {
	res, err := h.deps.App.Queries.ListClients.Handle(c.Request.Context(), query.FetchClientList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
