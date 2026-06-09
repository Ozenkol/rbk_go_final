package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/proposal"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProposalHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewProposalHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ProposalHandler {
	return &ProposalHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/proposals proposals createProposal
// Create a new proposal.
// Security:
//   Bearer:
// responses:
//   201: body:Proposal
//   400: body:errorResponse
func (h *ProposalHandler) CreateProposal(c *gin.Context) {
	var req http_requests.CreateProposalRequest
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

	p := &proposal.Proposal{
		ID:        uuid.New().String(),
		UserID:    userID,
		ClientID:  req.ClientID,
		DealID:    req.DealID,
		CompanyID: companyID,
		Title:     req.Title,
		Status:    shared.ProposalStatusDraft,
		Currency:  req.Currency,
		Discount:  req.Discount,
	}

	res, err := h.deps.App.Commands.CreateProposal.Handle(c.Request.Context(), command.CreateProposalCommand{Proposal: p})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/proposals/{id} proposals getProposal
// Get proposal by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Proposal
//   404: body:errorResponse
func (h *ProposalHandler) GetProposal(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetProposalByID.Handle(c.Request.Context(), query.FetchProposalByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/proposals/{id} proposals updateProposal
// Update an existing proposal.
// Security:
//   Bearer:
// responses:
//   200: body:Proposal
//   400: body:errorResponse
//   404: body:errorResponse
func (h *ProposalHandler) UpdateProposal(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateProposalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := h.deps.App.Queries.GetProposalByID.Handle(c.Request.Context(), query.FetchProposalByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proposal not found"})
		return
	}

	if req.Title != "" { p.Title = req.Title }
	if req.Status != "" { p.Status = shared.ProposalStatus(req.Status) }
	if req.Discount != 0 { p.Discount = req.Discount }

	res, err := h.deps.App.Commands.UpdateProposal.Handle(c.Request.Context(), command.UpdateProposalCommand{Proposal: p})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/proposals/{id} proposals deleteProposal
// Delete proposal by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *ProposalHandler) DeleteProposal(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteProposal.Handle(c.Request.Context(), command.DeleteProposalCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
