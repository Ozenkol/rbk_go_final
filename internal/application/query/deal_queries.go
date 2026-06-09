package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/deal"
)

type FetchDealByID struct {
	ID string
}

type FetchDealByIDHandler struct {
	repo deal.DealRepositoryInterface
}

func NewFetchDealByIDHandler(repo deal.DealRepositoryInterface) *FetchDealByIDHandler {
	return &FetchDealByIDHandler{repo: repo}
}

func (h *FetchDealByIDHandler) Handle(ctx context.Context, q FetchDealByID) (*deal.Deal, error) {
	return h.repo.GetByID(q.ID)
}

type FetchDealList struct{}

type FetchDealListHandler struct {
	repo deal.DealRepositoryInterface
}

func NewFetchDealListHandler(repo deal.DealRepositoryInterface) *FetchDealListHandler {
	return &FetchDealListHandler{repo: repo}
}

func (h *FetchDealListHandler) Handle(ctx context.Context, q FetchDealList) ([]*deal.Deal, error) {
	return h.repo.List()
}
