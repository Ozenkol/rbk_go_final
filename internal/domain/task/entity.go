package task

type Task struct {
	ID          string
	UserID    string
	CompanyID string
	Title       string
	Description string
	StartTime   string
	EndTime     string
	IsDone      bool
}
