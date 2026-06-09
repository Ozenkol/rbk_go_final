package proposal

type ProposalItem struct {
	ID          string  `json:"id"`
	ProposalID  string  `json:"proposal_id"`
	ProductID   string  `json:"product_id"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
