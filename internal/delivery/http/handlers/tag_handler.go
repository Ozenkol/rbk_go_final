package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/tag"
	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewTagHandler(deps *http_deps.Dependencies, logs *slog.Logger) *TagHandler {
	return &TagHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/tags tags createTag
// Создать новый тег.
// Security:
//   Bearer:
// responses:
//   201: getTagResponse
//   400: errorResponse
func (h *TagHandler) Create(c *gin.Context) {
	var req http_requests.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	_, companyID, err := h.deps.App.Services.AuthService.GetAuthInfoFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	t := &tag.Tag{
		CompanyID: companyID,
		Name:      req.Name,
	}

	res, err := h.deps.App.Commands.CreateTag.Handle(c.Request.Context(), command.CreateTagCommand{
		Tag: t,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/tags/{id} tags getTag
// Получить тег по ID.
// Security:
//   Bearer:
// responses:
//   200: getTagResponse
//   404: errorResponse
func (h *TagHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetTagByID.Handle(c.Request.Context(), query.FetchTagByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route GET /api/v1/tags tags listTags
// Список всех тегов.
// Security:
//   Bearer:
// responses:
//   200: []getTagResponse
//   500: errorResponse
func (h *TagHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListTags.Handle(c.Request.Context(), query.FetchTagList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/tags/{id} tags updateTag
// Обновить тег по ID.
// Security:
//   Bearer:
// responses:
//   200: getTagResponse
//   400: errorResponse
func (h *TagHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("Authorization")
	_, companyID, err := h.deps.App.Services.AuthService.GetAuthInfoFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	t := &tag.Tag{
		ID:        id,
		CompanyID: companyID,
		Name:      req.Name,
	}

	res, err := h.deps.App.Commands.UpdateTag.Handle(c.Request.Context(), command.UpdateTagCommand{
		Tag: t,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/tags/{id} tags deleteTag
// Удалить тег по ID.
// Security:
//   Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *TagHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteTag.Handle(c.Request.Context(), command.DeleteTagCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
