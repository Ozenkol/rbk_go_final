package persistence

import (
	"encoding/json"
	"github.com/Ozenkol/rbk-go-final/internal/domain/proposal"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type ProposalModel struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	UserID      string
	ClientID    string
	DealID      string
	CompanyID   string
	Title       string
	Status      string
	TotalAmount float64
	Currency    string
	Discount    float64
	ValidUntil  int64
	ItemsJSON   string `gorm:"type:text"`
}

type ProposalRepository struct {
	db *gorm.DB
}

func NewProposalRepository(db *gorm.DB) (proposal.ProposalRepositoryInterface, error) {
	if err := db.AutoMigrate(&ProposalModel{}); err != nil {
		return nil, err
	}
	return &ProposalRepository{db: db}, nil
}

func (r *ProposalRepository) Create(p *proposal.Proposal) (*proposal.Proposal, error) {
	model, err := toProposalModel(p)
	if err != nil {
		return nil, err
	}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return toProposalDomain(model)
}

func (r *ProposalRepository) GetByID(id string) (*proposal.Proposal, error) {
	var model ProposalModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return toProposalDomain(&model)
}

func (r *ProposalRepository) Update(p *proposal.Proposal) (*proposal.Proposal, error) {
	model, err := toProposalModel(p)
	if err != nil {
		return nil, err
	}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return toProposalDomain(model)
}

func (r *ProposalRepository) Delete(id string) error {
	return r.db.Delete(&ProposalModel{}, "id = ?", id).Error
}

func (r *ProposalRepository) List() ([]*proposal.Proposal, error) {
	var models []ProposalModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	proposals := make([]*proposal.Proposal, len(models))
	for i, m := range models {
		prop, err := toProposalDomain(&m)
		if err != nil {
			return nil, err
		}
		proposals[i] = prop
	}
	return proposals, nil
}

func toProposalModel(p *proposal.Proposal) (*ProposalModel, error) {
	itemsJSON, err := json.Marshal(p.ProposalItems)
	if err != nil {
		return nil, err
	}
	return &ProposalModel{
		ID:          p.ID,
		UserID:      p.UserID,
		ClientID:    p.ClientID,
		DealID:      p.DealID,
		CompanyID:   p.CompanyID,
		Title:       p.Title,
		Status:      string(p.Status),
		TotalAmount: p.TotalAmount,
		Currency:    p.Currency,
		Discount:    p.Discount,
		ValidUntil:  p.ValidUntil,
		ItemsJSON:   string(itemsJSON),
	}, nil
}

func toProposalDomain(m *ProposalModel) (*proposal.Proposal, error) {
	var items []proposal.ProposalItem
	if m.ItemsJSON != "" {
		if err := json.Unmarshal([]byte(m.ItemsJSON), &items); err != nil {
			return nil, err
		}
	}
	return &proposal.Proposal{
		ID:            m.ID,
		UserID:        m.UserID,
		ClientID:      m.ClientID,
		DealID:        m.DealID,
		CompanyID:     m.CompanyID,
		Title:         m.Title,
		Status:        shared.ProposalStatus(m.Status),
		TotalAmount:   m.TotalAmount,
		Currency:      m.Currency,
		Discount:      m.Discount,
		ValidUntil:    m.ValidUntil,
		CreatedAt:     m.CreatedAt.Unix(),
		UpdatedAt:     m.UpdatedAt.Unix(),
		ProposalItems: items,
	}, nil
}
