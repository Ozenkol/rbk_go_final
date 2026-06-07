package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
	"gorm.io/gorm"
)

type AnalyticModel struct {
	ID       string `gorm:"primaryKey"`
	ClientID string
	Data     string // Simplified for now
}

type AnalyticRepository struct {
	db *gorm.DB
}

func NewAnalyticRepository(db *gorm.DB) (analytic.AnalyticRepositoryInterface, error) {
	if err := db.AutoMigrate(&AnalyticModel{}); err != nil {
		return nil, err
	}
	return &AnalyticRepository{db: db}, nil
}

func (r *AnalyticRepository) Create(a *analytic.Analytic) (*analytic.Analytic, error) {
	model := &AnalyticModel{ID: a.ID, ClientID: a.ClientID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &analytic.Analytic{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *AnalyticRepository) GetByID(id string) (*analytic.Analytic, error) {
	var model AnalyticModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &analytic.Analytic{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *AnalyticRepository) Update(a *analytic.Analytic) (*analytic.Analytic, error) {
	model := &AnalyticModel{ID: a.ID, ClientID: a.ClientID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &analytic.Analytic{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *AnalyticRepository) Delete(id string) error {
	return r.db.Delete(&AnalyticModel{}, "id = ?", id).Error
}

func (r *AnalyticRepository) List() ([]*analytic.Analytic, error) {
	var models []AnalyticModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	analytics := make([]*analytic.Analytic, len(models))
	for i, m := range models {
		analytics[i] = &analytic.Analytic{ID: m.ID, ClientID: m.ClientID}
	}
	return analytics, nil
}
