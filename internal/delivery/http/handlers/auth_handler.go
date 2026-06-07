package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	"github.com/Ozenkol/rbk-go-final/internal/domain/user"
	"github.com/Ozenkol/rbk-go-final/internal/domain/shared"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewAuthHandler(deps *http_deps.Dependencies, logs *slog.Logger) *AuthHandler {
	return &AuthHandler{deps: deps, logs: logs}
}

// RegisterUser handles user registration
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req http_requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.deps.App.Commands.CreateUser.Handle(c.Request.Context(), command.CreateUserCommand{
		User: &user.User{
			HumanName: shared.HumanName{
				FirstName: req.FirstName,
				LastName:  req.LastName,
			},
			Email:    req.Email,
			Password: req.Password, // Command handler should handle hashing if following factory
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUserResponse(u))
}

// LoginUser handles user login
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req http_requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair, err := h.deps.App.Services.AuthService.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}

func createdUserResponse(u *user.User) gin.H {
	return gin.H{
		"id":    u.ID,
		"email": u.Email,
	}
}
