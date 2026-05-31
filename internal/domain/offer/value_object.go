package offer

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type OfferItem struct {
	Description string
	Amount      int
	ProductID   string
	Money       shared.Money
}