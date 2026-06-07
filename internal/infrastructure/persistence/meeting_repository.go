package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
	"gorm.io/gorm"
)

type MeetingModel struct {
	ID   string `gorm:"primaryKey"`
}

type MeetingRepository struct {
	db *gorm.DB
}

func NewMeetingRepository(db *gorm.DB) (meeting.MeetingRepositoryInterface, error) {
	if err := db.AutoMigrate(&MeetingModel{}); err != nil {
		return nil, err
	}
	return &MeetingRepository{db: db}, nil
}

func (r *MeetingRepository) Create(m *meeting.Meeting) (*meeting.Meeting, error) {
	model := &MeetingModel{ID: m.ID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &meeting.Meeting{ID: model.ID}, nil
}

func (r *MeetingRepository) GetByID(id string) (*meeting.Meeting, error) {
	var model MeetingModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &meeting.Meeting{ID: model.ID}, nil
}

func (r *MeetingRepository) Update(m *meeting.Meeting) (*meeting.Meeting, error) {
	model := &MeetingModel{ID: m.ID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &meeting.Meeting{ID: model.ID}, nil
}

func (r *MeetingRepository) Delete(id string) error {
	return r.db.Delete(&MeetingModel{}, "id = ?", id).Error
}

func (r *MeetingRepository) List() ([]*meeting.Meeting, error) {
	var models []MeetingModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	meetings := make([]*meeting.Meeting, len(models))
	for idx, m := range models {
		meetings[idx] = &meeting.Meeting{ID: m.ID}
	}
	return meetings, nil
}
