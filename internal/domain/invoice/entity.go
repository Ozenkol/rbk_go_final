package invoice

type Invoice struct {
	ID           string
	ClientID     string
	DocumentID   string
	CompanyID	string
	InvoiceItems []InvoiceItem
}

func (i *Invoice) TotalAmount() int {
	var total int
	for _, item := range i.InvoiceItems {
		total += item.Amount
	}
	return total
}

func (i *Invoice) AddItem(item InvoiceItem) {
	i.InvoiceItems = append(i.InvoiceItems, item)
}