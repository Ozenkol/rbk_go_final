package handlers

import (
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_types "github.com/Ozenkol/rbk-go-final/internal/delivery/http/types"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	deps *http_types.Dependencies
}

func NewUserHandler(deps *http_types.Dependencies) *UserHandler {
	return &UserHandler{deps: deps}
}

// @INVALID_ANNOTATION
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req command.CreateUserCommand
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// GetUser godoc
// @Summary Get user by ID
// @Description Fetch a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} user.User
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	q := query.FetchUserQuery{UserID: c.Param("id")}
	user, err := h.deps.App.Queries.GetUserByID.Handle(q)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}