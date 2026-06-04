package query

import "github.com/Ozenkol/rbk-go-final/internal/domain/task"

type FetchTaskByIDQuery struct {
	ID string
}

type FetchTaskByIDHandler struct {
	taskRepo task.TaskRepositoryInterface
}

func NewFetchTaskByIDHandler(taskRepo task.TaskRepositoryInterface) *FetchTaskByIDHandler {
	return &FetchTaskByIDHandler{
		taskRepo: taskRepo,
	}
}

func (h *FetchTaskByIDHandler) Handle(query FetchTaskByIDQuery) (*task.Task, error) {
	return h.taskRepo.GetByID(query.ID)
}