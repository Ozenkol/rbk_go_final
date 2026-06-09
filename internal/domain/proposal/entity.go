package proposal

import (
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
)

type Proposal struct {
	ID            string                `json:"id"`
	UserID        string                `json:"user_id"`
	CompanyID     string                `json:"company_id"`
	ClientID      string                `json:"client_id"`
	DealID        string                `json:"deal_id"`
	Title         string                `json:"title"`
	Status        shared.ProposalStatus `json:"status"`
	TotalAmount   float64               `json:"total_amount"`
	Currency      string                `json:"currency"`
	Discount      float64               `json:"discount"`
	ValidUntil    int64                 `json:"valid_until"`
	CreatedAt     int64                 `json:"created_at"`
	UpdatedAt     int64                 `json:"updated_at"`
	ProposalItems []ProposalItem        `json:"proposal_items"`
}

func (p *Proposal) CalculateTotal() float64 {
	var total float64
	for _, item := range p.ProposalItems {
		total += item.Price * float64(item.Quantity)
	}
	p.TotalAmount = total - p.Discount
	return p.TotalAmount
}
