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

type AuthHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewAuthHandler(deps *http_deps.Dependencies, logs *slog.Logger) *AuthHandler {
	return &AuthHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/auth/register users registerUser
//
// Register a new user and authenticate.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// responses:
//   201: createUserResponse
//   400: errorResponse
func (h *AuthHandler) RegisterUser(c *gin.Context) {
    var req http_requests.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        h.logs.Error("Invalid request body", slog.Any("error", err))
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // explicit mapping — delivery layer owns this translation
    cmd := command.CreateUserCommand{
        FirstName:  req.FirstName,
        MiddleName: req.MiddleName,
        LastName:   req.LastName,
        Email:      req.Email,
        Password:   req.Password,
    }

    id, err := h.deps.App.Commands.CreateUser.Handle(cmd)
    if err != nil {
        h.logs.Error("Failed to create user", slog.Any("error", err))
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err = h.deps.App.Queries.GetUserByID.Handle(query.FetchUserQuery{UserID: id})
    if err != nil {
        h.logs.Error("Failed to fetch user", slog.Any("error", err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    token, err := h.deps.App.Services.AuthService.Authenticate(cmd.Email, cmd.Password)
    if err != nil {
        h.logs.Error("Authentication failed", slog.Any("error", err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
        return
    }
    h.logs.Info("User registered and authenticated", slog.String("user_id", id))
    c.JSON(http.StatusCreated, 
        gin.H{
            "token": token,
        },
    )
}


// swagger:route POST /api/v1/auth/login users loginUser
//
// Authenticate an existing user.
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// responses:
//   200: loginUserResponse
//   400: errorResponse

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var req http_requests.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logs.Error("Invalid request body", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.deps.App.Services.AuthService.Authenticate(req.Email, req.Password)
	if err != nil {
		h.logs.Error("Authentication failed", slog.Any("error", err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	h.logs.Info("User logged in", slog.String("email", req.Email))
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// swagger:response loginUserResponse
type LoginUserResponse struct {
	// in: body
	Body struct {
		Token string `json:"token"`
	}
}

// swagger:parameters loginUser
type LoginUserRequestParams struct {
	// in: body
	// required: true
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}