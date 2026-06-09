package deal

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Deal struct {
	ID            string           `json:"id"`
	UserID        string           `json:"user_id"`
	CompanyID     string           `json:"company_id"`
	ClientID      string           `json:"client_id"`
	ResponsibleID string           `json:"responsible_id"`
	Title         string           `json:"title"`
	Stage         shared.DealStage `json:"stage"`
	Amount        float64          `json:"amount"`
	Currency      string           `json:"currency"`
	Probability   int              `json:"probability"` // 0-100
	Deadline      int64            `json:"deadline"`
	CreatedAt     int64            `json:"created_at"`
	UpdatedAt     int64            `json:"updated_at"`
}

func NewDeal(clientID string, title string, amount float64) Deal {
	return Deal{
		ClientID: clientID,
		Title:    title,
		Amount:   amount,
		Stage:    shared.DealStageNew,
	}
}
