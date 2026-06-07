package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type NotesFilter struct {
	ContentContains string
	CreatedAfter    string
	CreatedBefore   string
}

type FetchNotesListQuery struct {
	ClientID   string
	Filter     NotesFilter
	Pagination application_shared.Pagination
}

type NotesQueryRepositoryInterface interface {
	FindByFilter(clientID string, filter NotesFilter, pagination application_shared.Pagination) ([]*note.Note, error)
}

type FetchNotesListHandler struct {
	notesQueryRepository NotesQueryRepositoryInterface
}

func NewFetchNotesListHandler(notesRepo NotesQueryRepositoryInterface) *FetchNotesListHandler {
	return &FetchNotesListHandler{notesQueryRepository: notesRepo}
}

func (h *FetchNotesListHandler) Handle(query FetchNotesListQuery) ([]*note.Note, error) {
	notes, err := h.notesQueryRepository.FindByFilter(query.ClientID, query.Filter, query.Pagination)
	if err != nil {
		return nil, err
	}
	return notes, nil
}