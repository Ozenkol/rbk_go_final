package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewTaskHandler(deps *http_deps.Dependencies, logs *slog.Logger) *TaskHandler {
	return &TaskHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/tasks tasks createTask
// Create a new task.
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
	userID, err := h.deps.App.Services.AuthService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	t := &task.Task{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}

	res, err := h.deps.App.Commands.CreateTask.Handle(c.Request.Context(), command.CreateTaskCommand{
		Task: t,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/tasks/{id} tasks getTask
// Get task by ID.
// security:
//   Bearer:
// responses:
//   200: getTaskResponse
//   500: errorResponse
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
// Update a task.
// security:
//   Bearer:
// responses:
//   200: getTaskResponse
//   500: errorResponse
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	userID, err := h.deps.App.Services.AuthService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	t := &task.Task{
		ID:     id,
		UserID: userID,
	}
	if req.Title != nil {
		t.Title = *req.Title
	}
	if req.Description != nil {
		t.Description = *req.Description
	}
	if req.StartTime != nil {
		t.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		t.EndTime = *req.EndTime
	}
	if req.IsDone != nil {
		t.IsDone = *req.IsDone
	}

	res, err := h.deps.App.Commands.UpdateTask.Handle(c.Request.Context(), command.UpdateTaskCommand{
		Task: t,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/tasks/{id} tasks deleteTask
// Delete a task.
// security:
//   Bearer:
// responses:
//   200: deleteTaskResponse
//   500: errorResponse
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteTask.Handle(c.Request.Context(), command.DeleteTaskCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
