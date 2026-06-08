package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewProductHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ProductHandler {
	return &ProductHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/products products createProduct
func (h *ProductHandler) Create(c *gin.Context) {
	var req http_requests.CreateProductRequest
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

	p := &product.Product{
		UserID:      userID,
		Name:        req.Name,
		CompanyID:   req.CompanyID,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Price:       req.Price,
	}

	res, err := h.deps.App.Commands.CreateProduct.Handle(c.Request.Context(), command.CreateProductCommand{Product: p})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/products/{id} products getProduct
func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetProductByID.Handle(c.Request.Context(), query.FetchProductByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/products/{id} products updateProduct
func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateProductRequest
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

	p := &product.Product{
		ID:          id,
		UserID:      userID,
		Name:        req.Name,
		CompanyID:   req.CompanyID,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Price:       req.Price,
	}

	res, err := h.deps.App.Commands.UpdateProduct.Handle(c.Request.Context(), command.UpdateProductCommand{Product: p})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/products/{id} products deleteProduct
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteProduct.Handle(c.Request.Context(), command.DeleteProductCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/products products listProducts
func (h *ProductHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListProducts.Handle(c.Request.Context(), query.FetchProductList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
