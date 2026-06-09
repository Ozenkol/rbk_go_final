package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/deal"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type DealModel struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	UserID        string
	CompanyID     string
	ClientID      string
	ResponsibleID string
	Title         string
	Stage         string
	Amount        float64
	Currency      string
	Probability   int
	Deadline      int64
	CreatedAtUnix int64
	UpdatedAtUnix int64
}

type DealRepository struct {
	db *gorm.DB
}

func NewDealRepository(db *gorm.DB) (deal.DealRepositoryInterface, error) {
	if err := db.AutoMigrate(&DealModel{}); err != nil {
		return nil, err
	}
	return &DealRepository{db: db}, nil
}

func (r *DealRepository) Create(d *deal.Deal) (*deal.Deal, error) {
	model := toDealModel(d)
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toDealDomain(model), nil
}

func (r *DealRepository) GetByID(id string) (*deal.Deal, error) {
	var model DealModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toDealDomain(&model), nil
}

func (r *DealRepository) Update(d *deal.Deal) (*deal.Deal, error) {
	model := toDealModel(d)
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toDealDomain(model), nil
}

func (r *DealRepository) Delete(id string) error {
	return r.db.Delete(&DealModel{}, "id = ?", id).Error
}

func (r *DealRepository) List() ([]*deal.Deal, error) {
	var models []DealModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	deals := make([]*deal.Deal, len(models))
	for i, model := range models {
		deals[i] = toDealDomain(&model)
	}
	return deals, nil
}

func toDealDomain(m *DealModel) *deal.Deal {
	return &deal.Deal{
		ID:            m.ID,
		UserID:        m.UserID,
		CompanyID:     m.CompanyID,
		ClientID:      m.ClientID,
		ResponsibleID: m.ResponsibleID,
		Title:         m.Title,
		Stage:         shared.DealStage(m.Stage),
		Amount:        m.Amount,
		Currency:      m.Currency,
		Probability:   m.Probability,
		Deadline:      m.Deadline,
		CreatedAt:     m.CreatedAtUnix,
		UpdatedAt:     m.UpdatedAtUnix,
	}
}

func toDealModel(d *deal.Deal) *DealModel {
	return &DealModel{
		ID:            d.ID,
		UserID:        d.UserID,
		CompanyID:     d.CompanyID,
		ClientID:      d.ClientID,
		ResponsibleID: d.ResponsibleID,
		Title:         d.Title,
		Stage:         string(d.Stage),
		Amount:        d.Amount,
		Currency:      d.Currency,
		Probability:   d.Probability,
		Deadline:      d.Deadline,
		CreatedAtUnix: d.CreatedAt,
		UpdatedAtUnix: d.UpdatedAt,
	}
}
