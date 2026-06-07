package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
)

type FetchProductByID struct {
	ID string
}

type FetchProductByIDHandler struct {
	repo product.ProductRepositoryInterface
}

func NewFetchProductByIDHandler(repo product.ProductRepositoryInterface) *FetchProductByIDHandler {
	return &FetchProductByIDHandler{repo: repo}
}

func (h *FetchProductByIDHandler) Handle(ctx context.Context, q FetchProductByID) (*product.Product, error) {
	return h.repo.GetByID(q.ID)
}

type FetchProductList struct{}

type FetchProductListHandler struct {
	repo product.ProductRepositoryInterface
}

func NewFetchProductListHandler(repo product.ProductRepositoryInterface) *FetchProductListHandler {
	return &FetchProductListHandler{repo: repo}
}

func (h *FetchProductListHandler) Handle(ctx context.Context, q FetchProductList) ([]*product.Product, error) {
	return h.repo.List()
}
