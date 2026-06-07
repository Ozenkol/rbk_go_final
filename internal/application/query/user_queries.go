package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
)

type FetchUserByID struct {
	ID string
}

type FetchUserByIDHandler struct {
	repo user.UserRepositoryInterface
}

func NewFetchUserByIDHandler(repo user.UserRepositoryInterface) *FetchUserByIDHandler {
	return &FetchUserByIDHandler{repo: repo}
}

func (h *FetchUserByIDHandler) Handle(ctx context.Context, q FetchUserByID) (*user.User, error) {
	return h.repo.GetByID(q.ID)
}

type FetchUserList struct{}

type FetchUserListHandler struct {
	repo user.UserRepositoryInterface
}

func NewFetchUserListHandler(repo user.UserRepositoryInterface) *FetchUserListHandler {
	return &FetchUserListHandler{repo: repo}
}

func (h *FetchUserListHandler) Handle(ctx context.Context, q FetchUserList) ([]*user.User, error) {
	return h.repo.List()
}
