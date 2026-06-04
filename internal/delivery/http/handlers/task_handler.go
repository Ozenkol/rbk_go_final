package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
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
//
// Create a new task.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// responses:
//   201: createTaskResponse
//   400: errorResponse
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req http_requests.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd := command.CreateTaskCommand{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	}
	id, err := h.deps.App.Commands.CreateTask.Handle(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// swagger:route GET /api/v1/tasks/{id} tasks getTask
//
// Get task by ID.
//
// Produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     type: string
//
// responses:
//   200: getTaskResponse
//   500: errorResponse
func (h *TaskHandler) GetByID(c *gin.Context) {
	query := query.FetchTaskByIDQuery{
		ID: c.Param("id"),
	}
	task, err := h.deps.App.Queries.GetTaskByID.Handle(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// swagger:route PUT /api/v1/tasks/{id} tasks updateTask
//
// Update a task.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     type: string
//   + name: request
//     in: body
//     required: true
//     schema:
//       $ref: '#/definitions/UpdateTaskRequest'
//
// responses:
//   200: getTaskResponse
//   500: errorResponse
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var req http_requests.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd := command.UpdateTaskCommand{
		ID:          c.Param("id"),
		Title:       *req.Title,
		Description: *req.Description,
		StartTime:   *req.StartTime,
		EndTime:     *req.EndTime,
	}
	updatedTask, err := h.deps.App.Commands.UpdateTask.Handle(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// swagger:route DELETE /api/v1/tasks/{id} tasks deleteTask
//
// Delete a task.
//
// Produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     type: string
//
// responses:
//   200: deleteTaskResponse
//   500: errorResponse
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	err := h.deps.App.Commands.DeleteTask.Handle(command.DeleteTaskCommand{
		ID: c.Param("id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// swagger:route PUT /api/v1/tasks/{id}/toggle tasks toggleTask
//
// Toggle task completion.
//
// Produces:
// - application/json
//
// parameters:
//   + name: id
//     in: path
//     required: true
//     type: string
//
// responses:
//   200: getTaskResponse
//   500: errorResponse
func (h *TaskHandler) ToggleTask(c *gin.Context) {
	cmd:= command.UpdateTaskCommand{
		ID: c.Param("id"),
	}
	updatedTask, err := h.deps.App.Commands.UpdateTask.Handle(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// swagger:response createTaskResponse
type CreateTaskResponse struct {
	// in: body
	Body struct {
		ID string `json:"id"`
	}
}

// swagger:response getTaskResponse
type GetTaskResponse struct {
	// in: body
	Body struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
	}
}

// swagger:response deleteTaskResponse
type DeleteTaskResponse struct {
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// swagger:parameters updateTask
type UpdateTaskRequestParams struct {
	// in: body
	Body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		StartTime   string `json:"start_time"`
		EndTime     string `json:"end_time"`
	}
}