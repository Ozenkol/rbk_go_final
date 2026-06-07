package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/setting"
)

type CreateSettingCommand struct {
	Setting *setting.Setting
}

type CreateSettingHandler struct {
	repo setting.SettingRepositoryInterface
}

func NewCreateSettingHandler(repo setting.SettingRepositoryInterface) *CreateSettingHandler {
	return &CreateSettingHandler{repo: repo}
}

func (h *CreateSettingHandler) Handle(ctx context.Context, cmd CreateSettingCommand) (*setting.Setting, error) {
	return h.repo.Create(cmd.Setting)
}

type UpdateSettingCommand struct {
	Setting *setting.Setting
}

type UpdateSettingHandler struct {
	repo setting.SettingRepositoryInterface
}

func NewUpdateSettingHandler(repo setting.SettingRepositoryInterface) *UpdateSettingHandler {
	return &UpdateSettingHandler{repo: repo}
}

func (h *UpdateSettingHandler) Handle(ctx context.Context, cmd UpdateSettingCommand) (*setting.Setting, error) {
	return h.repo.Update(cmd.Setting)
}

type DeleteSettingCommand struct {
	ID string
}

type DeleteSettingHandler struct {
	repo setting.SettingRepositoryInterface
}

func NewDeleteSettingHandler(repo setting.SettingRepositoryInterface) *DeleteSettingHandler {
	return &DeleteSettingHandler{repo: repo}
}

func (h *DeleteSettingHandler) Handle(ctx context.Context, cmd DeleteSettingCommand) error {
	return h.repo.Delete(cmd.ID)
}
