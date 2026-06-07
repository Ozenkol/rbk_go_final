package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type FetchDocumentByID struct {
	ID string
}

type FetchDocumentByIDHandler struct {
	repo document.DocumentRepositoryInterface
}

func NewFetchDocumentByIDHandler(repo document.DocumentRepositoryInterface) *FetchDocumentByIDHandler {
	return &FetchDocumentByIDHandler{repo: repo}
}

func (h *FetchDocumentByIDHandler) Handle(ctx context.Context, q FetchDocumentByID) (*document.Document, error) {
	return h.repo.GetByID(q.ID)
}

type FetchDocumentList struct{}

type FetchDocumentListHandler struct {
	repo document.DocumentRepositoryInterface
}

func NewFetchDocumentListHandler(repo document.DocumentRepositoryInterface) *FetchDocumentListHandler {
	return &FetchDocumentListHandler{repo: repo}
}

func (h *FetchDocumentListHandler) Handle(ctx context.Context, q FetchDocumentList) ([]*document.Document, error) {
	return h.repo.List()
}
