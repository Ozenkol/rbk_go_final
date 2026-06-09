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
	"github.com/google/uuid"
)

type ProductHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewProductHandler(deps *http_deps.Dependencies, logs *slog.Logger) *ProductHandler {
	return &ProductHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/products products createProduct
// Create a new product.
// Security:
//   Bearer:
// responses:
//   201: body:Product
//   400: body:errorResponse
func (h *ProductHandler) Create(c *gin.Context) {
	var req http_requests.CreateProductRequest
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

	prod := &product.Product{
		ID:          uuid.New().String(),
		UserID:      userID,
		CompanyID:   companyID,
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Category:    req.Category,
		SKU:         req.SKU,
		Price:       req.Price,
		Currency:    req.Currency,
		Unit:        req.Unit,
		IsActive:    req.IsActive,
	}

	res, err := h.deps.App.Commands.CreateProduct.Handle(c.Request.Context(), command.CreateProductCommand{Product: prod})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/products/{id} products getProduct
// Get product by ID.
// Security:
//   Bearer:
// responses:
//   200: body:Product
//   404: body:errorResponse
func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetProductByID.Handle(c.Request.Context(), query.FetchProductByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/products products listProducts
// List all products.
// Security:
//   Bearer:
// responses:
//   200: body:[]Product
func (h *ProductHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListProducts.Handle(c.Request.Context(), query.FetchProductList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/products/{id} products updateProduct
// Update an existing product.
// Security:
//   Bearer:
// responses:
//   200: body:Product
//   400: body:errorResponse
//   404: body:errorResponse
func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod, err := h.deps.App.Queries.GetProductByID.Handle(c.Request.Context(), query.FetchProductByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if req.Name != "" { prod.Name = req.Name }
	if req.Description != "" { prod.Description = req.Description }
	if req.Price != 0 { prod.Price = req.Price }
	prod.IsActive = req.IsActive

	res, err := h.deps.App.Commands.UpdateProduct.Handle(c.Request.Context(), command.UpdateProductCommand{Product: prod})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/products/{id} products deleteProduct
// Delete product by ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: body:errorResponse
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteProduct.Handle(c.Request.Context(), command.DeleteProductCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
