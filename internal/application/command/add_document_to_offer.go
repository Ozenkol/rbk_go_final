package command

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/Ozenkol/rbk-go-final/internal/domain/offer"
)

type AddDocuemntToOfferCommand struct {
	OfferID    string
	DocumentID string
}

type AddDocumentToOfferHandler struct {
	offerRepo offer.OfferRepositoryInterface
	documentRepo document.DocumentRepositoryInterface
}

func NewAddDocumentToOfferHandler(offerRepo offer.OfferRepositoryInterface, documentRepo document.DocumentRepositoryInterface) *AddDocumentToOfferHandler {
	return &AddDocumentToOfferHandler{
		offerRepo: offerRepo,
		documentRepo: documentRepo,
	}
}

func (h *AddDocumentToOfferHandler) Handle(cmd AddDocuemntToOfferCommand) error {
	offer, err := h.offerRepo.GetByID(cmd.OfferID)
	if err != nil {
		return err
	}

	document, err := h.documentRepo.GetByID(cmd.DocumentID)
	if err != nil {
		return err
	}

	offer.AddDocument(*document)

	return h.offerRepo.Update(offer)
}