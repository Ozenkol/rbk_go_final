package task

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Task struct {
	ID            string              `json:"id"`
	UserID        string              `json:"user_id"`
	CompanyID     string              `json:"company_id"`
	ClientID      string              `json:"client_id"`
	DealID        string              `json:"deal_id"`
	ContractID    string              `json:"contract_id"`
	ResponsibleID string              `json:"responsible_id"`
	Title         string              `json:"title"`
	Description   string              `json:"description"`
	Status        shared.TaskStatus   `json:"status"`
	Priority      shared.TaskPriority `json:"priority"`
	Deadline      int64               `json:"deadline"`
	ReminderAt    int64               `json:"reminder_at"`
	CreatedAt     int64               `json:"created_at"`
	UpdatedAt     int64               `json:"updated_at"`
}
