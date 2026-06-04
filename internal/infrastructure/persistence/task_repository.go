package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"gorm.io/gorm"
)

type TaskModel struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    string
	Title       string
	Description string
	StartTime   string
	EndTime     string
	IsDone      bool
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) (task.TaskRepositoryInterface, error) {
	if err := db.AutoMigrate(&TaskModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &TaskRepository{db: db}, nil
}

func (r *TaskRepository) Create(task *task.Task) (*task.Task, error) {
	taskModel := toTaskModel(task)
	err := r.db.Create(taskModel).Error
	if err != nil {
		return nil, err
	}
	return toTaskDomain(taskModel), nil
}

func (r *TaskRepository) GetByID(id string) (*task.Task, error) {
	var model TaskModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return toTaskDomain(&model), nil
}

func (r *TaskRepository) Update(task *task.Task) (*task.Task, error) {
	taskModel := toTaskModel(task)
	err := r.db.Save(taskModel).Error
	if err != nil {
		return nil, err
	}
	return toTaskDomain(taskModel), nil
}

func (r *TaskRepository) Delete(id string) error {
	return r.db.Delete(&TaskModel{}, "id = ?", id).Error
}

func toTaskModel(task *task.Task) *TaskModel {
	taskModel := &TaskModel{
		UserID:    task.UserID,
		Title:       task.Title,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		IsDone:      task.IsDone,
	}
	if task.ID != "" {
		taskModel.ID = task.ID
	}
	return taskModel
}

func toTaskDomain(model *TaskModel) *task.Task {
	return &task.Task{
		ID:          model.ID,
		UserID:    model.UserID,
		Title:       model.Title,
		Description: model.Description,
		StartTime:   model.StartTime,
		EndTime:     model.EndTime,
		IsDone:      model.IsDone,
	}
}
