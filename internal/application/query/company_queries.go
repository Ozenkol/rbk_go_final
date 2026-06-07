package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
)

type FetchCompanyByID struct {
	ID string
}

type FetchCompanyByIDHandler struct {
	repo company.CompanyRepositoryInterface
}

func NewFetchCompanyByIDHandler(repo company.CompanyRepositoryInterface) *FetchCompanyByIDHandler {
	return &FetchCompanyByIDHandler{repo: repo}
}

func (h *FetchCompanyByIDHandler) Handle(ctx context.Context, q FetchCompanyByID) (*company.Company, error) {
	return h.repo.GetByID(q.ID)
}

type FetchCompanyList struct{}

type FetchCompanyListHandler struct {
	repo company.CompanyRepositoryInterface
}

func NewFetchCompanyListHandler(repo company.CompanyRepositoryInterface) *FetchCompanyListHandler {
	return &FetchCompanyListHandler{repo: repo}
}

func (h *FetchCompanyListHandler) Handle(ctx context.Context, q FetchCompanyList) ([]*company.Company, error) {
	return h.repo.List()
}
