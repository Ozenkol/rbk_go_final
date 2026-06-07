package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
)

type FetchAnalyticByID struct {
	ID string
}

type FetchAnalyticByIDHandler struct {
	repo analytic.AnalyticRepositoryInterface
}

func NewFetchAnalyticByIDHandler(repo analytic.AnalyticRepositoryInterface) *FetchAnalyticByIDHandler {
	return &FetchAnalyticByIDHandler{repo: repo}
}

func (h *FetchAnalyticByIDHandler) Handle(ctx context.Context, q FetchAnalyticByID) (*analytic.Analytic, error) {
	return h.repo.GetByID(q.ID)
}

type FetchAnalyticList struct{}

type FetchAnalyticListHandler struct {
	repo analytic.AnalyticRepositoryInterface
}

func NewFetchAnalyticListHandler(repo analytic.AnalyticRepositoryInterface) *FetchAnalyticListHandler {
	return &FetchAnalyticListHandler{repo: repo}
}

func (h *FetchAnalyticListHandler) Handle(ctx context.Context, q FetchAnalyticList) ([]*analytic.Analytic, error) {
	return h.repo.List()
}
