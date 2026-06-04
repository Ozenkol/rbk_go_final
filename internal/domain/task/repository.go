package task

type TaskRepositoryInterface interface {
	Create(task *Task) (*Task, error)
	GetByID(id string) (*Task, error)
	Update(task *Task) (*Task, error)
	Delete(id string) error
}
