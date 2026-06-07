package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type DeleteNoteCommand struct {
	ID string
}

type DeleteNoteHandler struct {
	repo note.NoteRepositoryInterface
}

func NewDeleteNoteHandler(repo note.NoteRepositoryInterface) *DeleteNoteHandler {
	return &DeleteNoteHandler{repo: repo}
}

func (h *DeleteNoteHandler) Handle(ctx context.Context, cmd DeleteNoteCommand) error {
	return h.repo.Delete(cmd.ID)
}
