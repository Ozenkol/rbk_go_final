package http_requests

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	StartTime   *string `json:"start_time,omitempty"`
	EndTime     *string `json:"end_time,omitempty"`
	IsDone      *bool   `json:"is_done,omitempty"`
}
