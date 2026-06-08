package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/file"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewFileHandler(deps *http_deps.Dependencies, logs *slog.Logger) *FileHandler {
	return &FileHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/files files createFile
func (h *FileHandler) Create(c *gin.Context) {
	var req http_requests.CreateFileRequest
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

	f := &file.File{
		UserID:           userID,
		ClientID:         req.ClientID,
		CompanyID:        req.CompanyID,
		StorageReference: req.StorageReference,
		Name:             req.Name,
	}

	res, err := h.deps.App.Commands.CreateFile.Handle(c.Request.Context(), command.CreateFileCommand{File: f})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/files/{id} files getFile
func (h *FileHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetFileByID.Handle(c.Request.Context(), query.FetchFileByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/files/{id} files updateFile
func (h *FileHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.UpdateFileRequest
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

	f := &file.File{
		ID:               id,
		UserID:           userID,
		ClientID:         req.ClientID,
		CompanyID:        req.CompanyID,
		StorageReference: req.StorageReference,
		Name:             req.Name,
	}

	res, err := h.deps.App.Commands.UpdateFile.Handle(c.Request.Context(), command.UpdateFileCommand{File: f})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/files/{id} files deleteFile
func (h *FileHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteFile.Handle(c.Request.Context(), command.DeleteFileCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// swagger:route GET /api/v1/files files listFiles
func (h *FileHandler) List(c *gin.Context) {
	res, err := h.deps.App.Queries.ListFiles.Handle(c.Request.Context(), query.FetchFileList{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
