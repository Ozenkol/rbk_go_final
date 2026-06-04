package handlers

import (
	"log/slog"

	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewNoteHandler(deps *http_deps.Dependencies, logs *slog.Logger) *NoteHandler {
	return &NoteHandler{deps: deps, logs: logs}
}	

func (h *NoteHandler) CreateNote(c *gin.Context) {	
	c.JSON(200, gin.H{"message": "Create note - protected route"})
}

func (h *NoteHandler) GetNote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get note - protected route"})
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update note - protected route"})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete note - protected route"})
}

