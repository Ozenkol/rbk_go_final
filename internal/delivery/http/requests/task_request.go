package http_requests

// CreateTaskRequest represents the request body for creating a task
// swagger:model CreateTaskRequest
type CreateTaskRequest struct {
	// Task title
	// example: Do offer
	Title       string `json:"title"`
	// Task description
	// example: Prepare and send offer to client
	Description string `json:"description"`
	// Task start time
	// example: 2023-10-27T10:00:00Z
	StartTime   string `json:"start_time"`
	// Task end time
	// example: 2023-10-27T11:00:00Z
	EndTime     string `json:"end_time"`
}

// UpdateTaskRequest represents the request body for updating a task
// swagger:model UpdateTaskRequest
type UpdateTaskRequest struct {
	// Task title
	// example: Do offer
	Title       *string `json:"title,omitempty"`
	// Task description
	// example: Prepare and send offer to client
	Description *string `json:"description,omitempty"`
	// Task start time
	// example: 2023-10-27T10:00:00Z
	StartTime   *string `json:"start_time,omitempty"`
	// Task end time
	// example: 2023-10-27T11:00:00Z
	EndTime     *string `json:"end_time,omitempty"`
	// Task completion status
	// example: true
	IsDone      *bool   `json:"is_done,omitempty"`
}
