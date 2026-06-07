package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/notification"
	"gorm.io/gorm"
)

type NotificationModel struct {
	ID       string `gorm:"primaryKey"`
	ClientID string
	Message  string
}

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) (notification.NotificationRepositoryInterface, error) {
	if err := db.AutoMigrate(&NotificationModel{}); err != nil {
		return nil, err
	}
	return &NotificationRepository{db: db}, nil
}

func (r *NotificationRepository) Create(n *notification.Notification) (*notification.Notification, error) {
	model := toNotificationModel(n)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toNotificationDomain(model), nil
}

func (r *NotificationRepository) GetByID(id string) (*notification.Notification, error) {
	var model NotificationModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toNotificationDomain(&model), nil
}

func (r *NotificationRepository) Update(n *notification.Notification) (*notification.Notification, error) {
	model := toNotificationModel(n)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toNotificationDomain(model), nil
}

func (r *NotificationRepository) Delete(id string) error {
	return r.db.Delete(&NotificationModel{}, "id = ?", id).Error
}

func (r *NotificationRepository) List() ([]*notification.Notification, error) {
	var models []NotificationModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	notifications := make([]*notification.Notification, len(models))
	for i, m := range models {
		notifications[i] = toNotificationDomain(&m)
	}
	return notifications, nil
}

func toNotificationModel(n *notification.Notification) *NotificationModel {
	return &NotificationModel{
		ID:       n.ID,
		ClientID: n.ClientID,
		Message:  n.Message,
	}
}

func toNotificationDomain(m *NotificationModel) *notification.Notification {
	return &notification.Notification{
		ID:       m.ID,
		ClientID: m.ClientID,
		Message:  m.Message,
	}
}
