package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewCompanyHandler(deps *http_deps.Dependencies, logs *slog.Logger) *CompanyHandler {
	return &CompanyHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/companies companies createCompany
func (h *CompanyHandler) Create(c *gin.Context) {
	var req http_requests.CreateCompanyRequest
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

	comp := &company.Company{
		UserID: userID,
		Name:   req.Name,
	}

	res, err := h.deps.App.Commands.CreateCompany.Handle(c.Request.Context(), command.CreateCompanyCommand{
		Company: comp,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/companies/{id} companies getCompany
func (h *CompanyHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetCompanyByID.Handle(c.Request.Context(), query.FetchCompanyByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/companies/{id} companies updateCompany
func (h *CompanyHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateCompanyRequest
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

	comp := &company.Company{
		ID:     id,
		UserID: userID,
		Name:   req.Name,
	}

	res, err := h.deps.App.Commands.UpdateCompany.Handle(c.Request.Context(), command.UpdateCompanyCommand{
		Company: comp,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/companies/{id} companies deleteCompany
func (h *CompanyHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteCompany.Handle(c.Request.Context(), command.DeleteCompanyCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/companies companies listCompanies
func (h *CompanyHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListCompanies.Handle(c.Request.Context(), query.FetchCompanyList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
