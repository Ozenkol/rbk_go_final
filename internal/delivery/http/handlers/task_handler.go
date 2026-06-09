package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewTaskHandler(deps *http_deps.Dependencies, logs *slog.Logger) *TaskHandler {
	return &TaskHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/tasks tasks createTask
// Создать новую задачу.
// Security:
//   Bearer:
// responses:
//   201: createTaskResponse
//   400: errorResponse
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req http_requests.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	userID, companyID, err := h.deps.App.Services.AuthService.GetAuthInfoFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	t := &task.Task{
		ID:            uuid.New().String(),
		UserID:        userID,
		CompanyID:     companyID,
		ClientID:      req.ClientID,
		DealID:        req.DealID,
		ContractID:    req.ContractID,
		ResponsibleID: req.ResponsibleID,
		Title:         req.Title,
		Description:   req.Description,
		Status:        shared.TaskStatusNew,
		Priority:      shared.TaskPriorityMedium,
		Deadline:      req.Deadline,
	}

	res, err := h.deps.App.Commands.CreateTask.Handle(c.Request.Context(), command.CreateTaskCommand{Task: t})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/tasks/{id} tasks getTask
// Получить задачу по ID.
// Security:
//   Bearer:
// responses:
//   200: getTaskResponse
//   404: errorResponse
func (h *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetTaskByID.Handle(c.Request.Context(), query.FetchTaskByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/tasks/{id} tasks updateTask
// Обновить задачу по ID.
// Security:
//   Bearer:
// responses:
//   200: getTaskResponse
//   400: errorResponse
//   404: errorResponse
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := h.deps.App.Queries.GetTaskByID.Handle(c.Request.Context(), query.FetchTaskByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if req.Title != "" { t.Title = req.Title }
	if req.Description != "" { t.Description = req.Description }
	if req.Status != "" { t.Status = shared.TaskStatus(req.Status) }
	if req.Priority != "" { t.Priority = shared.TaskPriority(req.Priority) }
	if req.Deadline != 0 { t.Deadline = req.Deadline }

	res, err := h.deps.App.Commands.UpdateTask.Handle(c.Request.Context(), command.UpdateTaskCommand{Task: t})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/tasks/{id} tasks deleteTask
// Удалить задачу по ID.
// Security:
//   Bearer:
// responses:
//   204: deleteTaskResponse
//   500: errorResponse
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteTask.Handle(c.Request.Context(), command.DeleteTaskCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
