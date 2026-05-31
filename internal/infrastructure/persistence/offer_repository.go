package persistence

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"gorm.io/gorm"
)

type OfferModel struct {
	ID         string `gorm:"primaryKey"`
	ClientID   string
	DocumentID string
	OfferItems []OfferItemModel `gorm:"foreignKey:OfferID"`
}

type OfferItemModel struct {
	ID uint `gorm:"primaryKey"`

	OfferID   string
	ProductID string

	Quantity int
	Price    float64
	Currency CurrencyModel `gorm:"foreignKey:Currency"`
}

type CurrencyModel struct {
	Code string
	Name string
}

type OfferRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) (offer.OfferRepositoryInterface, error) {
	if err := db.AutoMigrate(&OfferModel{}, &OfferItemModel{}, &CurrencyModel{}); err != nil {
		panic(err) // Handle error properly in production
	}
	return &OfferRepository{db: db}, nil
}

func (r *OfferRepository) Create(offer *offer.Offer) error {
	offerModel := toOfferModel(offer)
	return r.db.Create(offerModel).Error
}

func (r *OfferRepository) GetByID(id string) (*offer.Offer, error) {
	var model OfferModel
	if err := r.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return toOfferDomain(&model), nil
}

func (r *OfferRepository) Update(offer *offer.Offer) error {
	offerModel := toOfferModel(offer)
	return r.db.Save(offerModel).Error
}

func (r *OfferRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&OfferModel{}).Error
}

func toOfferDomain(offerModel *OfferModel) *offer.Offer {
	offerItems := make([]offer.OfferItem, len(offerModel.OfferItems))
	for i, item := range offerModel.OfferItems {
		offerItems[i] = offer.OfferItem{
			Description: "",
			Amount:      item.Quantity,
			ProductID:   item.ProductID,
			Money:       shared.Money{Amount: item.Price, Currency: shared.Currency{Code: item.Currency.Code, Name: item.Currency.Name}},
		}
	}
	return &offer.Offer{
		ID:         offerModel.ID,
		ClientID:   offerModel.ClientID,
		DocumentID: offerModel.DocumentID,
		OfferItems: offerItems,
	}
}

func toOfferModel(offer *offer.Offer) *OfferModel {
	offerItems := make([]OfferItemModel, len(offer.OfferItems))
	for i, item := range offer.OfferItems {
		currency := CurrencyModel{
			Code: item.Money.Currency.Code,
			Name: item.Money.Currency.Name,
		}
		offerItems[i] = OfferItemModel{
			OfferID:   offer.ID,
			ProductID: item.ProductID,
			Quantity:  int(item.Amount),
			Price:     item.Money.Amount,
			Currency:  currency,
		}
	}
	return &OfferModel{
		ID:         offer.ID,
		ClientID:   offer.ClientID,
		DocumentID: offer.DocumentID,
		OfferItems: offerItems,
	}
}
