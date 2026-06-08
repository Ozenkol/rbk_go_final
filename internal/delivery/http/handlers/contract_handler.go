package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
	"github.com/gin-gonic/gin"
)

type ContractHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewContractHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ContractHandler {
	return &ContractHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/contracts contracts createContract
func (h *ContractHandler) Create(c *gin.Context) {
	var req http_requests.CreateContractRequest
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

	cont := &contract.Contract{
		UserID:    userID,
		ClientID:  req.ClientID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
	}

	res, err := h.deps.App.Commands.CreateContract.Handle(c.Request.Context(), command.CreateContractCommand{
		Contract: cont,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/contracts/{id} contracts getContract
func (h *ContractHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetContractByID.Handle(c.Request.Context(), query.FetchContractByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/contracts/{id} contracts updateContract
func (h *ContractHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateContractRequest
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

	cont := &contract.Contract{
		ID:        id,
		UserID:    userID,
		ClientID:  req.ClientID,
		CompanyID: req.CompanyID,
		Content:   req.Content,
	}

	res, err := h.deps.App.Commands.UpdateContract.Handle(c.Request.Context(), command.UpdateContractCommand{
		Contract: cont,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/contracts/{id} contracts deleteContract
func (h *ContractHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteContract.Handle(c.Request.Context(), command.DeleteContractCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/contracts contracts listContracts
func (h *ContractHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListContracts.Handle(c.Request.Context(), query.FetchContractList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
