package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
)

type TaskFilter struct {
	isDone *bool
	ClientID string
}

type TaskQueryRespositoryInterface interface {
	GetByID(id string) (*task.Task, error)
	FindByFilter(filter TaskFilter, pagination application_shared.Pagination) ([]*task.Task, error)
}

type FetchTaskListQuery struct {
	Filter TaskFilter
	Pagination application_shared.Pagination
}

type FetchTaskListHandler struct {
	taskQueryRepository TaskQueryRespositoryInterface
}

func NewFetchTaskListHandler(taskRepo TaskQueryRespositoryInterface) *FetchTaskListHandler {
	return &FetchTaskListHandler{taskQueryRepository: taskRepo}
}

func (h *FetchTaskListHandler) Handle(query FetchTaskListQuery) ([]*task.Task, error) {
	tasks, err := h.taskQueryRepository.FindByFilter(query.Filter, query.Pagination)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}