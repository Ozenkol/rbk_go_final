package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type CreateNoteCommand struct {
	Note *note.Note
}

type CreateNoteHandler struct {
	repo note.NoteRepositoryInterface
}

func NewCreateNoteHandler(repo note.NoteRepositoryInterface) *CreateNoteHandler {
	return &CreateNoteHandler{repo: repo}
}

func (h *CreateNoteHandler) Handle(ctx context.Context, cmd CreateNoteCommand) (*note.Note, error) {
	return h.repo.Create(cmd.Note)
}

type UpdateNoteCommand struct {
	Note *note.Note
}

type UpdateNoteHandler struct {
	repo note.NoteRepositoryInterface
}

func NewUpdateNoteHandler(repo note.NoteRepositoryInterface) *UpdateNoteHandler {
	return &UpdateNoteHandler{repo: repo}
}

func (h *UpdateNoteHandler) Handle(ctx context.Context, cmd UpdateNoteCommand) (*note.Note, error) {
	return h.repo.Update(cmd.Note)
}
