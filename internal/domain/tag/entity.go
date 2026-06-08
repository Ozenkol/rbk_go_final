package tag

type Tag struct {
	ID     string `json:"id"`
	// swagger:ignore
	UserID string `json:"user_id"`
	// swagger:ignore
	CompanyID string `json:"company_id"`
	Name   string `json:"name"`
}
