package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContractHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewContractHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ContractHandler {
	return &ContractHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/contracts contracts createContract
// Create a new contract.
// Security:
//   Bearer:
// responses:
//   201: body:Contract
//   400: body:errorResponse
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

	con := &contract.Contract{
		ID:         uuid.New().String(),
		UserID:     userID,
		ClientID:   req.ClientID,
		DealID:     req.DealID,
		CompanyID:  companyID,
		Number:     req.Number,
		Status:     shared.ContractStatus(req.Status),
		ValidFrom:  req.ValidFrom,
		ValidUntil: req.ValidUntil,
	}

	res, err := h.deps.App.Commands.CreateContract.Handle(c.Request.Context(), command.CreateContractCommand{Contract: con})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/contracts/{id} contracts getContract
// Get contract by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Contract
//   404: body:errorResponse
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
// List all contracts.
// Security:
//   Bearer:
// responses:
//   200: body:[]Contract
func (h *ContractHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListContracts.Handle(c.Request.Context(), query.FetchContractList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/contracts/{id} contracts updateContract
// Update an existing contract.
// Security:
//   Bearer:
// responses:
//   200: body:Contract
//   400: body:errorResponse
//   404: body:errorResponse
func (h *ContractHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateContractRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resByID, err := h.deps.App.Queries.GetContractByID.Handle(c.Request.Context(), query.FetchContractByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}
	con := resByID
	if req.Number != "" { con.Number = req.Number }
	if req.Status != "" { con.Status = shared.ContractStatus(req.Status) }
	if req.ValidFrom != 0 { con.ValidFrom = req.ValidFrom }
	if req.ValidUntil != 0 { con.ValidUntil = req.ValidUntil }

	res, err := h.deps.App.Commands.UpdateContract.Handle(c.Request.Context(), command.UpdateContractCommand{Contract: con})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/contracts/{id} contracts deleteContract
// Delete contract by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *ContractHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteContract.Handle(c.Request.Context(), command.DeleteContractCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
