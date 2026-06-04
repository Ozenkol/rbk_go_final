package command

import "github.com/Ozenkol/rbk-go-final/internal/domain/task"

type CreateTaskCommand struct {
	UserID      string
	Title       string
	Description string
	StartTime   string
	EndTime     string
}

type CreateTaskHandler struct {
	taskRepo task.TaskRepositoryInterface
}

func NewCreateTaskHandler(taskRepo task.TaskRepositoryInterface) *CreateTaskHandler {
	return &CreateTaskHandler{taskRepo: taskRepo}
}

func (h *CreateTaskHandler) Handle(cmd CreateTaskCommand) (string, error) {
	task := &task.Task{
		UserID:    cmd.UserID,
		Title:       cmd.Title,
		Description: cmd.Description,
		StartTime:   cmd.StartTime,
		EndTime:     cmd.EndTime,
	}
	savedTask, err := h.taskRepo.Create(task)
	if err != nil {
		return "", err
	}
	return savedTask.ID, nil
}