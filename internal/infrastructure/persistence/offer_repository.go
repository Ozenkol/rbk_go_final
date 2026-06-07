package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"gorm.io/gorm"
)

type OfferModel struct {
	ID       string `gorm:"primaryKey"`
	ClientID string
}

type OfferRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) (offer.OfferRepositoryInterface, error) {
	if err := db.AutoMigrate(&OfferModel{}); err != nil {
		return nil, err
	}
	return &OfferRepository{db: db}, nil
}

func (r *OfferRepository) Create(o *offer.Offer) (*offer.Offer, error) {
	model := &OfferModel{ID: o.ID, ClientID: o.ClientID}
	if err := r.db.Create(model).Error; err != nil {
		return nil, err
	}
	return &offer.Offer{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *OfferRepository) GetByID(id string) (*offer.Offer, error) {
	var model OfferModel
	if err := r.db.First(&model, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &offer.Offer{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *OfferRepository) Update(o *offer.Offer) (*offer.Offer, error) {
	model := &OfferModel{ID: o.ID, ClientID: o.ClientID}
	if err := r.db.Save(model).Error; err != nil {
		return nil, err
	}
	return &offer.Offer{ID: model.ID, ClientID: model.ClientID}, nil
}

func (r *OfferRepository) Delete(id string) error {
	return r.db.Delete(&OfferModel{}, "id = ?", id).Error
}

func (r *OfferRepository) List() ([]*offer.Offer, error) {
	var models []OfferModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	offers := make([]*offer.Offer, len(models))
	for i, m := range models {
		offers[i] = &offer.Offer{ID: m.ID, ClientID: m.ClientID}
	}
	return offers, nil
}
