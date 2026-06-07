package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
)

type FetchSettingByID struct {
	ID string
}

type FetchSettingByIDHandler struct {
	repo setting.SettingRepositoryInterface
}

func NewFetchSettingByIDHandler(repo setting.SettingRepositoryInterface) *FetchSettingByIDHandler {
	return &FetchSettingByIDHandler{repo: repo}
}

func (h *FetchSettingByIDHandler) Handle(ctx context.Context, q FetchSettingByID) (*setting.Setting, error) {
	return h.repo.GetByID(q.ID)
}

type FetchSettingList struct{}

type FetchSettingListHandler struct {
	repo setting.SettingRepositoryInterface
}

func NewFetchSettingListHandler(repo setting.SettingRepositoryInterface) *FetchSettingListHandler {
	return &FetchSettingListHandler{repo: repo}
}

func (h *FetchSettingListHandler) Handle(ctx context.Context, q FetchSettingList) ([]*setting.Setting, error) {
	return h.repo.List()
}
