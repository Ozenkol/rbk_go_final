package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
	"gorm.io/gorm"
)

type CommunicationModel struct {
	ID   string `gorm:"primaryKey"`
	Type string
}

type CommunicationRepository struct {
	db *gorm.DB
}

func NewCommunicationRepository(db *gorm.DB) (communication.CommunicationRepositoryInterface, error) {
	if err := db.AutoMigrate(&CommunicationModel{}); err != nil {
		return nil, err
	}
	return &CommunicationRepository{db: db}, nil
}

func (r *CommunicationRepository) Create(c *communication.Communication) (*communication.Communication, error) {
	model := &CommunicationModel{ID: c.ID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &communication.Communication{ID: model.ID}, nil
}

func (r *CommunicationRepository) GetByID(id string) (*communication.Communication, error) {
	var model CommunicationModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &communication.Communication{ID: model.ID}, nil
}

func (r *CommunicationRepository) Update(c *communication.Communication) (*communication.Communication, error) {
	model := &CommunicationModel{ID: c.ID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &communication.Communication{ID: model.ID}, nil
}

func (r *CommunicationRepository) Delete(id string) error {
	return r.db.Delete(&CommunicationModel{}, "id = ?", id).Error
}

func (r *CommunicationRepository) List() ([]*communication.Communication, error) {
	var models []CommunicationModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	comms := make([]*communication.Communication, len(models))
	for i, m := range models {
		comms[i] = &communication.Communication{ID: m.ID}
	}
	return comms, nil
}
