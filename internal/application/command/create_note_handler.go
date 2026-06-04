package command

import "github.com/Ozenkol/rbk-go-final/internal/domain/note"

type CreateNoteCommand struct {
	UserID   string
	ClientID string
	Content  string
}

type CreateNoteHandler struct {
	noteRepo note.NoteRepositoryInterface
}

func NewCreateNoteHandler(noteRepo note.NoteRepositoryInterface) *CreateNoteHandler {
	return &CreateNoteHandler{noteRepo: noteRepo}
}

func (h *CreateNoteHandler) Handle(cmd CreateNoteCommand) (string, error) {
	note := &note.Note{
		UserID:   cmd.UserID,
		ClientID: cmd.ClientID,
		Content:  cmd.Content,
	}
	savedNote, err := h.noteRepo.Create(note)
	if err != nil {
		return "", err
	}
	return savedNote.ID, nil
}