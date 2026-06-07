package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
)

type FetchFileByID struct {
	ID string
}

type FetchFileByIDHandler struct {
	repo file.FileRepositoryInterface
}

func NewFetchFileByIDHandler(repo file.FileRepositoryInterface) *FetchFileByIDHandler {
	return &FetchFileByIDHandler{repo: repo}
}

func (h *FetchFileByIDHandler) Handle(ctx context.Context, q FetchFileByID) (*file.File, error) {
	return h.repo.GetByID(q.ID)
}

type FetchFileList struct{}

type FetchFileListHandler struct {
	repo file.FileRepositoryInterface
}

func NewFetchFileListHandler(repo file.FileRepositoryInterface) *FetchFileListHandler {
	return &FetchFileListHandler{repo: repo}
}

func (h *FetchFileListHandler) Handle(ctx context.Context, q FetchFileList) ([]*file.File, error) {
	return h.repo.List()
}
