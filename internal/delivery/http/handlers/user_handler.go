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

type UserHandler struct {
	deps *http_deps.Dependencies
    logs *slog.Logger
}

func NewUserHandler(deps *http_deps.Dependencies, logs *slog.Logger) *UserHandler {
	return &UserHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/users users createUser
//
// Create a new user.
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
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req http_requests.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.deps.App.Queries.GetUserByID.Handle(query.FetchUserQuery{UserID: id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, 
        gin.H{
            "id": user.ID,
            "first_name": user.HumanName.FirstName,
            "middle_name": user.HumanName.MiddleName,
            "last_name": user.HumanName.LastName,
            "email": user.Email,
        },
    )

}

// swagger:route GET /api/v1/users/{id} users getUser
//
// Get user by ID.
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
//   200: getUserResponse
//   404: errorResponse
func (h *UserHandler) GetUser(c *gin.Context) {
	q := query.FetchUserQuery{UserID: c.Param("id")}
	user, err := h.deps.App.Queries.GetUserByID.Handle(q)
	if err != nil {
		h.logs.Error("User not found", slog.String("user_id", c.Param("id")))
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	h.logs.Info("User retrieved", slog.String("user_id", c.Param("id")))
	c.JSON(http.StatusOK, user)
}

// swagger:response createUserResponse
type CreateUserResponse struct {
	// in: body
	Body struct {
		ID         string `json:"id"`
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
	}
}

// swagger:response getUserResponse
type GetUserResponse struct {
	// in: body
	Body struct {
		ID         string `json:"id"`
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
	}
}

// swagger:response errorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		Error string `json:"error"`
	}
}

// swagger:parameters createUser
type CreateUserRequestParams struct {
	// in: body
	// required: true
	Body struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}
}
