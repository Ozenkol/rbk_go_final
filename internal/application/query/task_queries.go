package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
)

type FetchTaskByID struct {
	ID string
}

type FetchTaskByIDHandler struct {
	repo task.TaskRepositoryInterface
}

func NewFetchTaskByIDHandler(repo task.TaskRepositoryInterface) *FetchTaskByIDHandler {
	return &FetchTaskByIDHandler{repo: repo}
}

func (h *FetchTaskByIDHandler) Handle(ctx context.Context, q FetchTaskByID) (*task.Task, error) {
	return h.repo.GetByID(q.ID)
}

type FetchTaskList struct{}

type FetchTaskListHandler struct {
	repo task.TaskRepositoryInterface
}

func NewFetchTaskListHandler(repo task.TaskRepositoryInterface) *FetchTaskListHandler {
	return &FetchTaskListHandler{repo: repo}
}

func (h *FetchTaskListHandler) Handle(ctx context.Context, q FetchTaskList) ([]*task.Task, error) {
	return h.repo.List()
}
