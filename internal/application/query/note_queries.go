package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type FetchNoteByID struct {
	ID string
}

type FetchNoteByIDHandler struct {
	repo note.NoteRepositoryInterface
}

func NewFetchNoteByIDHandler(repo note.NoteRepositoryInterface) *FetchNoteByIDHandler {
	return &FetchNoteByIDHandler{repo: repo}
}

func (h *FetchNoteByIDHandler) Handle(ctx context.Context, q FetchNoteByID) (*note.Note, error) {
	return h.repo.GetByID(q.ID)
}

type FetchNoteList struct{}

type FetchNoteListHandler struct {
	repo note.NoteRepositoryInterface
}

func NewFetchNoteListHandler(repo note.NoteRepositoryInterface) *FetchNoteListHandler {
	return &FetchNoteListHandler{repo: repo}
}

func (h *FetchNoteListHandler) Handle(ctx context.Context, q FetchNoteList) ([]*note.Note, error) {
	return h.repo.List()
}
