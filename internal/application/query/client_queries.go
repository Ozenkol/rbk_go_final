package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
)

type FetchClientByID struct {
	ID string
}

type FetchClientByIDHandler struct {
	repo client.ClientRepositoryInterface
}

func NewFetchClientByIDHandler(repo client.ClientRepositoryInterface) *FetchClientByIDHandler {
	return &FetchClientByIDHandler{repo: repo}
}

func (h *FetchClientByIDHandler) Handle(ctx context.Context, q FetchClientByID) (*client.Client, error) {
	return h.repo.GetByID(q.ID)
}

type FetchClientList struct{}

type FetchClientListHandler struct {
	repo client.ClientRepositoryInterface
}

func NewFetchClientListHandler(repo client.ClientRepositoryInterface) *FetchClientListHandler {
	return &FetchClientListHandler{repo: repo}
}

func (h *FetchClientListHandler) Handle(ctx context.Context, q FetchClientList) ([]*client.Client, error) {
	return h.repo.List()
}
