package query

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type FetchNoteById struct {
	ID string
}

type FetchNoteByIdHandler struct {
	noteRepo note.NoteRepositoryInterface
}

func NewFetchNoteByIdHandler(noteRepo note.NoteRepositoryInterface) *FetchNoteByIdHandler {
	return &FetchNoteByIdHandler{
		noteRepo: noteRepo,
	}
}

func (h *FetchNoteByIdHandler) Handle(query FetchNoteById) (*note.Note, error) {
	return h.noteRepo.GetByID(query.ID)
}