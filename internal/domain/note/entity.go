package note

type Note struct {
	ID       string `json:"id"`
	// swagger:ignore
	UserID   string `json:"user_id"`
	ClientID string `json:"client_id"`
	// swagger:ignore
	CompanyID string `json:"company_id"`
	Content  string `json:"content"`
}
