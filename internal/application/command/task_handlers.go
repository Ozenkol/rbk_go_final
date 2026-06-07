package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
)

type CreateTaskCommand struct {
	Task *task.Task
}

type CreateTaskHandler struct {
	repo task.TaskRepositoryInterface
}

func NewCreateTaskHandler(repo task.TaskRepositoryInterface) *CreateTaskHandler {
	return &CreateTaskHandler{repo: repo}
}

func (h *CreateTaskHandler) Handle(ctx context.Context, cmd CreateTaskCommand) (*task.Task, error) {
	return h.repo.Create(cmd.Task)
}

type UpdateTaskCommand struct {
	Task *task.Task
}

type UpdateTaskHandler struct {
	repo task.TaskRepositoryInterface
}

func NewUpdateTaskHandler(repo task.TaskRepositoryInterface) *UpdateTaskHandler {
	return &UpdateTaskHandler{repo: repo}
}

func (h *UpdateTaskHandler) Handle(ctx context.Context, cmd UpdateTaskCommand) (*task.Task, error) {
	return h.repo.Update(cmd.Task)
}

type DeleteTaskCommand struct {
	ID string
}

type DeleteTaskHandler struct {
	repo task.TaskRepositoryInterface
}

func NewDeleteTaskHandler(repo task.TaskRepositoryInterface) *DeleteTaskHandler {
	return &DeleteTaskHandler{repo: repo}
}

func (h *DeleteTaskHandler) Handle(ctx context.Context, cmd DeleteTaskCommand) error {
	return h.repo.Delete(cmd.ID)
}
