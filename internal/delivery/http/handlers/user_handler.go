package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewUserHandler(deps *http_deps.Dependencies, logs *slog.Logger) *UserHandler {
	return &UserHandler{deps: deps, logs: logs}
}

// swagger:route GET /api/v1/users/{id} users getUser
// Get a user by ID.
// Security:
//   Bearer:
// responses:
//   200: getUserResponse
//   404: errorResponse
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	u, err := h.deps.App.Queries.GetUserByID.Handle(c.Request.Context(), query.FetchUserByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// swagger:route POST /api/v1/users users createUser
// CreateUser (for consistency, though registration is preferred)
func (h *UserHandler) CreateUser(c *gin.Context) {
	// Re-use Auth logic or similar
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Use /auth/register"})
}
