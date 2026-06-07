package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
)

type CreateTagCommand struct {
	Tag *tag.Tag
}

type CreateTagHandler struct {
	repo tag.TagRepositoryInterface
}

func NewCreateTagHandler(repo tag.TagRepositoryInterface) *CreateTagHandler {
	return &CreateTagHandler{repo: repo}
}

func (h *CreateTagHandler) Handle(ctx context.Context, cmd CreateTagCommand) (*tag.Tag, error) {
	return h.repo.Create(cmd.Tag)
}

type UpdateTagCommand struct {
	Tag *tag.Tag
}

type UpdateTagHandler struct {
	repo tag.TagRepositoryInterface
}

func NewUpdateTagHandler(repo tag.TagRepositoryInterface) *UpdateTagHandler {
	return &UpdateTagHandler{repo: repo}
}

func (h *UpdateTagHandler) Handle(ctx context.Context, cmd UpdateTagCommand) (*tag.Tag, error) {
	return h.repo.Update(cmd.Tag)
}

type DeleteTagCommand struct {
	ID string
}

type DeleteTagHandler struct {
	repo tag.TagRepositoryInterface
}

func NewDeleteTagHandler(repo tag.TagRepositoryInterface) *DeleteTagHandler {
	return &DeleteTagHandler{repo: repo}
}

func (h *DeleteTagHandler) Handle(ctx context.Context, cmd DeleteTagCommand) error {
	return h.repo.Delete(cmd.ID)
}
