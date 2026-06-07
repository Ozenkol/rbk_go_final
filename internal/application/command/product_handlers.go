package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
)

type CreateProductCommand struct {
	Product *product.Product
}

type CreateProductHandler struct {
	repo product.ProductRepositoryInterface
}

func NewCreateProductHandler(repo product.ProductRepositoryInterface) *CreateProductHandler {
	return &CreateProductHandler{repo: repo}
}

func (h *CreateProductHandler) Handle(ctx context.Context, cmd CreateProductCommand) (*product.Product, error) {
	return h.repo.Create(cmd.Product)
}

type UpdateProductCommand struct {
	Product *product.Product
}

type UpdateProductHandler struct {
	repo product.ProductRepositoryInterface
}

func NewUpdateProductHandler(repo product.ProductRepositoryInterface) *UpdateProductHandler {
	return &UpdateProductHandler{repo: repo}
}

func (h *UpdateProductHandler) Handle(ctx context.Context, cmd UpdateProductCommand) (*product.Product, error) {
	return h.repo.Update(cmd.Product)
}

type DeleteProductCommand struct {
	ID string
}

type DeleteProductHandler struct {
	repo product.ProductRepositoryInterface
}

func NewDeleteProductHandler(repo product.ProductRepositoryInterface) *DeleteProductHandler {
	return &DeleteProductHandler{repo: repo}
}

func (h *DeleteProductHandler) Handle(ctx context.Context, cmd DeleteProductCommand) error {
	return h.repo.Delete(cmd.ID)
}
