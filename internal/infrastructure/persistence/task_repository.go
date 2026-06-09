package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"gorm.io/gorm"
)

type TaskModel struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	UserID        string
	CompanyID     string
	ClientID      string
	DealID        string
	ContractID    string
	ResponsibleID string
	Title         string
	Description   string
	Status        string
	Priority      string
	Deadline      int64
	ReminderAt    int64
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) (task.TaskRepositoryInterface, error) {
	if err := db.AutoMigrate(&TaskModel{}); err != nil {
		return nil, err
	}
	return &TaskRepository{db: db}, nil
}

func (r *TaskRepository) Create(t *task.Task) (*task.Task, error) {
	model := toTaskModel(t)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toTaskDomain(model), nil
}

func (r *TaskRepository) GetByID(id string) (*task.Task, error) {
	var model TaskModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toTaskDomain(&model), nil
}

func (r *TaskRepository) Update(t *task.Task) (*task.Task, error) {
	model := toTaskModel(t)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toTaskDomain(model), nil
}

func (r *TaskRepository) Delete(id string) error {
	return r.db.Delete(&TaskModel{}, "id = ?", id).Error
}

func (r *TaskRepository) List() ([]*task.Task, error) {
	var models []TaskModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	tasks := make([]*task.Task, len(models))
	for i, m := range models {
		tasks[i] = toTaskDomain(&m)
	}
	return tasks, nil
}

func toTaskModel(t *task.Task) *TaskModel {
	return &TaskModel{
		ID:            t.ID,
		UserID:        t.UserID,
		CompanyID:     t.CompanyID,
		ClientID:      t.ClientID,
		DealID:        t.DealID,
		ContractID:    t.ContractID,
		ResponsibleID: t.ResponsibleID,
		Title:         t.Title,
		Description:   t.Description,
		Status:        string(t.Status),
		Priority:      string(t.Priority),
		Deadline:      t.Deadline,
		ReminderAt:    t.ReminderAt,
	}
}

func toTaskDomain(m *TaskModel) *task.Task {
	return &task.Task{
		ID:            m.ID,
		UserID:        m.UserID,
		CompanyID:     m.CompanyID,
		ClientID:      m.ClientID,
		DealID:        m.DealID,
		ContractID:    m.ContractID,
		ResponsibleID: m.ResponsibleID,
		Title:         m.Title,
		Description:   m.Description,
		Status:        shared.TaskStatus(m.Status),
		Priority:      shared.TaskPriority(m.Priority),
		Deadline:      m.Deadline,
		ReminderAt:    m.ReminderAt,
		CreatedAt:     m.CreatedAt.Unix(),
		UpdatedAt:     m.UpdatedAt.Unix(),
	}
}
