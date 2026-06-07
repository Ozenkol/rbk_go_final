package command

import (
	"errors"

	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
)

type UpdateNoteCommand struct {
	ID          string
	UserID      string
	Title       string
	Description string
	StartTime   string
	EndTime     string
}

type UpdateNoteHandler struct {
	noteRepo note.NoteRepositoryInterface
}

func NewUpdateNoteHandler(noteRepo note.NoteRepositoryInterface) *UpdateNoteHandler {
	return &UpdateNoteHandler{noteRepo: noteRepo}
}

func (h *UpdateTaskHandler) Handle(cmd UpdateTaskCommand) (string, error) {
	existingTask, err := h.taskRepo.GetByID(cmd.ID)
	if err != nil {
		return "", err
	}
	if existingTask == nil {
		return "", errors.New("task not found")
	}
	if cmd.UserID != existingTask.UserID {
		return "", errors.New("user does not have permission to update this task")
	}
	if cmd.Title != "" {
		existingTask.Title = cmd.Title
	}
	if cmd.Description != "" {
		existingTask.Description = cmd.Description
	}
	if cmd.StartTime != "" {
		existingTask.StartTime = cmd.StartTime
	}
	if cmd.EndTime != "" {
		existingTask.EndTime = cmd.EndTime
	}
	
	savedTask, err := h.taskRepo.Update(existingTask)
	if err != nil {
		return "", err
	}
	return savedTask.ID, nil
}