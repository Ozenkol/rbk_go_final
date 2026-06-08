package task

type Task struct {
	ID          string `json:"id"`
	// swagger:ignore
	UserID    string `json:"user_id"`
	// swagger:ignore
	CompanyID string `json:"company_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	IsDone      bool   `json:"is_done"`
}
