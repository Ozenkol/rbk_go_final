package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"gorm.io/gorm"
)

type TaskModel struct {
	ID          string `gorm:"primaryKey"`
	ClientID    string
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

func (r *TaskRepository) Create(task *task.Task) error {
	taskModel := toTaskModel(task)
	return r.db.Create(taskModel).Error
}

func (r *TaskRepository) GetByID(id string) (*task.Task, error) {
	var model TaskModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return toTaskDomain(&model), nil
}

func (r *TaskRepository) Update(task *task.Task) error {
	taskModel := toTaskModel(task)
	return r.db.Save(taskModel).Error
}

func (r *TaskRepository) Delete(id string) error {
	return r.db.Delete(&TaskModel{}, "id = ?", id).Error
}

func toTaskModel(task *task.Task) *TaskModel {
	return &TaskModel{
		ID:          task.ID,
		ClientID:    task.ClientID,
		Title:       task.Title,
		Description: task.Description,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
		IsDone:      task.IsDone,
	}
}


func toTaskDomain(model *TaskModel) *task.Task {
	return &task.Task{
		ID:          model.ID,
		ClientID:    model.ClientID,
		Title:       model.Title,
		Description: model.Description,
		StartTime:   model.StartTime,
		EndTime:     model.EndTime,
		IsDone:      model.IsDone,
	}
}