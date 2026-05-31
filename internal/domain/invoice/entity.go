package invoice

type Invoice struct {
	ID           string
	ClientID     string
	DocumentID   string
	InvoiceItems []InvoiceItem
}
