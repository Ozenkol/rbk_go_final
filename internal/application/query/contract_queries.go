package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
)

type FetchContractByID struct {
	ID string
}

type FetchContractByIDHandler struct {
	repo contract.ContractRepositoryInterface
}

func NewFetchContractByIDHandler(repo contract.ContractRepositoryInterface) *FetchContractByIDHandler {
	return &FetchContractByIDHandler{repo: repo}
}

func (h *FetchContractByIDHandler) Handle(ctx context.Context, q FetchContractByID) (*contract.Contract, error) {
	return h.repo.GetByID(q.ID)
}

type FetchContractList struct{}

type FetchContractListHandler struct {
	repo contract.ContractRepositoryInterface
}

func NewFetchContractListHandler(repo contract.ContractRepositoryInterface) *FetchContractListHandler {
	return &FetchContractListHandler{repo: repo}
}

func (h *FetchContractListHandler) Handle(ctx context.Context, q FetchContractList) ([]*contract.Contract, error) {
	return h.repo.List()
}
