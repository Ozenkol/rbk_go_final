package command

import (
	"errors"

	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type UpdateNoteCommand struct {
	ID          string
	UserID      string
	ClientID	string
	Content     string
}

type UpdateNoteHandler struct {
	noteRepo note.NoteRepositoryInterface
}

func NewUpdateNoteHandler(noteRepo note.NoteRepositoryInterface) *UpdateNoteHandler {
	return &UpdateNoteHandler{noteRepo: noteRepo}
}

func (h *UpdateNoteHandler) Handle(cmd UpdateNoteCommand) (string, error) {
	existingNote, err := h.noteRepo.GetByID(cmd.ID)
	if err != nil {
		return "", err
	}
	if existingNote == nil {
		return "", errors.New("note not found")
	}
	if cmd.UserID != existingNote.UserID {
		return "", errors.New("user does not have permission to update this note")
	}
	if cmd.Content != "" {
		existingNote.Content = cmd.Content
	}
	
	savedNote, err := h.noteRepo.Update(existingNote)
	if err != nil {
		return "", err
	}
	return savedNote.ID, nil
}