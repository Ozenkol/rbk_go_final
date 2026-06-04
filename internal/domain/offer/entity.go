package offer

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
	"github.com/Ozenkol/rbk-go-final/internal/domain/product"
)

type Offer struct {
	ID         string
	ClientID   string
	DocumentID string
	OfferItems []OfferItem
}

func (o *Offer) TotalAmount() float64 {
	var total float64
	for _, item := range o.OfferItems {
		total += item.Money.Amount * float64(item.Amount)
	}
	return total
}

func (o *Offer) AddOfferItem(product product.Product) {
	isExistItem, foundOffer := findOfferItem(o.OfferItems, product.ID)
	if isExistItem {
		o.OfferItems = append(o.OfferItems, OfferItem{
			ProductID: product.ID,
			Amount:    1,
			Money:     product.Price,
		})
	} else {
		foundOffer.Amount++
	}
}

func (o *Offer) DecrementOfferItem(product product.Product, amount int) {
	for i, item := range o.OfferItems {
		if item.ProductID == product.ID {
			o.OfferItems[i].Amount -= amount
			if o.OfferItems[i].Amount <= 0 {
				o.RemoveOfferItem(product)
			}
			break
		}
	}
}

func (o *Offer) RemoveOfferItem(product product.Product) {
	for i, item := range o.OfferItems {
		if item.ProductID == product.ID {
			o.OfferItems = append(o.OfferItems[:i], o.OfferItems[i+1:]...)
			break
		}
	}
}

func (o *Offer) AddDocument(document document.Document) {
	if o.DocumentID != "" {
		return
	}
	o.DocumentID = document.ID
}

func findOfferItem(offerItems []OfferItem, productID string) (bool, *OfferItem) {
	for _, item := range offerItems {
		if item.ProductID == productID {
			return true, &item
		}
	}
	return false, nil
}