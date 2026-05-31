package task

type TaskRepositoryInterface interface {
	Create(task *Task) error
	GetByID(id string) (*Task, error)
	Update(task *Task) error
	Delete(id string) error
}
