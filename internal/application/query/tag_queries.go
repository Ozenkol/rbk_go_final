package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
)

type FetchTagByID struct {
	ID string
}

type FetchTagByIDHandler struct {
	repo tag.TagRepositoryInterface
}

func NewFetchTagByIDHandler(repo tag.TagRepositoryInterface) *FetchTagByIDHandler {
	return &FetchTagByIDHandler{repo: repo}
}

func (h *FetchTagByIDHandler) Handle(ctx context.Context, q FetchTagByID) (*tag.Tag, error) {
	return h.repo.GetByID(q.ID)
}

type FetchTagList struct{}

type FetchTagListHandler struct {
	repo tag.TagRepositoryInterface
}

func NewFetchTagListHandler(repo tag.TagRepositoryInterface) *FetchTagListHandler {
	return &FetchTagListHandler{repo: repo}
}

func (h *FetchTagListHandler) Handle(ctx context.Context, q FetchTagList) ([]*tag.Tag, error) {
	return h.repo.List()
}
