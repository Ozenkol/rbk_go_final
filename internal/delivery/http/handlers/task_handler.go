package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/task"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
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
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.deps.App.Commands.CreateTask.Handle(c.Request.Context(), command.CreateTaskCommand{Task: &t})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/tasks/{id} tasks getTask
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
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id
	res, err := h.deps.App.Commands.UpdateTask.Handle(c.Request.Context(), command.UpdateTaskCommand{Task: &t})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/tasks/{id} tasks deleteTask
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteTask.Handle(c.Request.Context(), command.DeleteTaskCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
