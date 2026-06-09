package http_requests

type CreateTaskRequest struct {
	ClientID      string `json:"client_id"`
	DealID        string `json:"deal_id"`
	ContractID    string `json:"contract_id"`
	ResponsibleID string `json:"responsible_id"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description"`
	Deadline      int64  `json:"deadline"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Deadline    int64  `json:"deadline"`
}
