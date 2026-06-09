package invoice

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Invoice struct {
	ID           string               `json:"id"`
	UserID       string               `json:"user_id"`
	ClientID     string               `json:"client_id"`
	DealID       string               `json:"deal_id"`
	CompanyID    string               `json:"company_id"`
	Number       string               `json:"number"`
	Status       shared.InvoiceStatus `json:"status"`
	TotalAmount  float64              `json:"total_amount"`
	PaidAmount   float64              `json:"paid_amount"`
	DueDate      int64                `json:"due_date"`
	CreatedAt    int64                `json:"created_at"`
	UpdatedAt    int64                `json:"updated_at"`
	InvoiceItems []InvoiceItem        `json:"invoice_items"`
}

func (i *Invoice) CalculateTotal() float64 {
	var total float64
	for _, item := range i.InvoiceItems {
		total += item.Price * float64(item.Quantity)
	}
	i.TotalAmount = total
	return total
}

func (i *Invoice) AddItem(item InvoiceItem) {
	i.InvoiceItems = append(i.InvoiceItems, item)
}