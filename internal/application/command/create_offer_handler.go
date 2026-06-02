package command

import "github.com/Ozenkol/rbk-go-final/internal/domain/offer"

type CreateOfferCommand struct {
	UserID      string   `json:"user_id"`
	ClientID    string   `json:"client_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ProductIDs  []string `json:"product_ids"`
}

type CreateOfferHandler struct {
	offerFactory    offer.OfferFactoryInterface
	offerRepository offer.OfferRepositoryInterface
}

func NewCreateOfferHandler(offerFactory offer.OfferFactoryInterface, offerRepository offer.OfferRepositoryInterface) *CreateOfferHandler {
	return &CreateOfferHandler{
		offerFactory:    offerFactory,
		offerRepository: offerRepository,
	}
}

func (h *CreateOfferHandler) Handle(cmd CreateOfferCommand) (string, error) {
	offer, err := h.offerFactory.CreateOffer(cmd.UserID, cmd.ClientID, cmd.Title, cmd.Description, cmd.ProductIDs)
	if err != nil {
		return "", err
	}
	err = h.offerRepository.Create(offer)
	if err != nil {
		return "", err
	}
	return offer.ID, nil
}