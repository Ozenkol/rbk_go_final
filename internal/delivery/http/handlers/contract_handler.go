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
// Создать новый контракт.
// Security:
//   Bearer:
// responses:
//   201: getContractResponse
//   400: errorResponse
func (h *ContractHandler) Create(c *gin.Context) {
	var req http_requests.CreateContractRequest
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

	cn := &contract.Contract{
		UserID:    userID,
		CompanyID: companyID,
		ClientID:  req.ClientID,
		Content:   req.Content,
	}

	res, err := h.deps.App.Commands.CreateContract.Handle(c.Request.Context(), command.CreateContractCommand{
		Contract: cn,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/contracts/{id} contracts getContract
// Получить контракт по ID.
// Security:
//   Bearer:
// responses:
//   200: getContractResponse
//   404: errorResponse
func (h *ContractHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetContractByID.Handle(c.Request.Context(), query.FetchContractByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/contracts contracts listContracts
// Список всех контрактов.
// Security:
//   Bearer:
// responses:
//   200: []getContractResponse
//   500: errorResponse
func (h *ContractHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListContracts.Handle(c.Request.Context(), query.FetchContractList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/contracts/{id} contracts updateContract
// Обновить контракт по ID.
// Security:
//   Bearer:
// responses:
//   200: getContractResponse
//   400: errorResponse
func (h *ContractHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateContractRequest
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

	cn := &contract.Contract{
		ID:        id,
		UserID:    userID,
		CompanyID: companyID,
		ClientID:  req.ClientID,
		Content:   req.Content,
	}

	res, err := h.deps.App.Commands.UpdateContract.Handle(c.Request.Context(), command.UpdateContractCommand{
		Contract: cn,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/contracts/{id} contracts deleteContract
// Удалить контракт по ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *ContractHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteContract.Handle(c.Request.Context(), command.DeleteContractCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
