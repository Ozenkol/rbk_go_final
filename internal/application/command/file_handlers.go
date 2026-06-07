package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
)

type CreateFileCommand struct {
	File *file.File
}

type CreateFileHandler struct {
	repo file.FileRepositoryInterface
}

func NewCreateFileHandler(repo file.FileRepositoryInterface) *CreateFileHandler {
	return &CreateFileHandler{repo: repo}
}

func (h *CreateFileHandler) Handle(ctx context.Context, cmd CreateFileCommand) (*file.File, error) {
	return h.repo.Create(cmd.File)
}

type UpdateFileCommand struct {
	File *file.File
}

type UpdateFileHandler struct {
	repo file.FileRepositoryInterface
}

func NewUpdateFileHandler(repo file.FileRepositoryInterface) *UpdateFileHandler {
	return &UpdateFileHandler{repo: repo}
}

func (h *UpdateFileHandler) Handle(ctx context.Context, cmd UpdateFileCommand) (*file.File, error) {
	return h.repo.Update(cmd.File)
}

type DeleteFileCommand struct {
	ID string
}

type DeleteFileHandler struct {
	repo file.FileRepositoryInterface
}

func NewDeleteFileHandler(repo file.FileRepositoryInterface) *DeleteFileHandler {
	return &DeleteFileHandler{repo: repo}
}

func (h *DeleteFileHandler) Handle(ctx context.Context, cmd DeleteFileCommand) error {
	return h.repo.Delete(cmd.ID)
}
