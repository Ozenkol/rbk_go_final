package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type CommunicationModel struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	UserID    string
	CompanyID string
	ClientID  string
	DealID    string
	Type      string
	Content   string
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
	model := toCommunicationModel(c)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toCommunicationDomain(model), nil
}

func (r *CommunicationRepository) GetByID(id string) (*communication.Communication, error) {
	var model CommunicationModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toCommunicationDomain(&model), nil
}

func (r *CommunicationRepository) Update(c *communication.Communication) (*communication.Communication, error) {
	model := toCommunicationModel(c)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toCommunicationDomain(model), nil
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
		comms[i] = toCommunicationDomain(&m)
	}
	return comms, nil
}

func toCommunicationModel(c *communication.Communication) *CommunicationModel {
	return &CommunicationModel{
		ID:        c.ID,
		UserID:    c.UserID,
		CompanyID: c.CompanyID,
		ClientID:  c.ClientID,
		DealID:    c.DealID,
		Type:      string(c.Type),
		Content:   c.Content,
	}
}

func toCommunicationDomain(m *CommunicationModel) *communication.Communication {
	return &communication.Communication{
		ID:        m.ID,
		UserID:    m.UserID,
		CompanyID: m.CompanyID,
		ClientID:  m.ClientID,
		DealID:    m.DealID,
		Type:      shared.CommunicationType(m.Type),
		Content:   m.Content,
		CreatedAt: m.CreatedAt.Unix(),
	}
}
