package command

import "github.com/Ozenkol/rbk-go-final/internal/domain/task"

type DeleteTaskCommand struct {
	ID string
}

type DeleteTaskHandler struct {
	taskRepo task.TaskRepositoryInterface
}

func NewDeleteTaskHandler(taskRepo task.TaskRepositoryInterface) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		taskRepo: taskRepo,
	}
}

func (h *DeleteTaskHandler) Handle(cmd DeleteTaskCommand) error {
	return h.taskRepo.Delete(cmd.ID)
}