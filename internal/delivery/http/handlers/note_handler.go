package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Ozenkol/rbk-go-final/internal/application/command"
	"github.com/Ozenkol/rbk-go-final/internal/application/query"
	http_deps "github.com/Ozenkol/rbk-go-final/internal/delivery/http/deps"
	http_requests "github.com/Ozenkol/rbk-go-final/internal/delivery/http/requests"
	"github.com/Ozenkol/rbk-go-final/internal/domain/note"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewNoteHandler(deps *http_deps.Dependencies, logs *slog.Logger) *NoteHandler {
	return &NoteHandler{deps: deps, logs: logs}
}

// swagger:route POST /api/v1/notes notes createNote
// Create a new note.
// security:
//   - Bearer:
// responses:
//   201: getNoteResponse
//   400: errorResponse
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req http_requests.CreateNoteRequest
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

	n := &note.Note{
		UserID:   userID,
		ClientID: req.ClientID,
		Content:  req.Content,
	}

	res, err := h.deps.App.Commands.CreateNote.Handle(c.Request.Context(), command.CreateNoteCommand{
		Note: n,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// swagger:route GET /api/v1/notes/{id} notes getNote
// Get a note by ID.
// security:
//   - Bearer:
// responses:
//   200: getNoteResponse
//   404: errorResponse
func (h *NoteHandler) GetNote(c *gin.Context) {
	id := c.Param("id")
	res, err := h.deps.App.Queries.GetNoteByID.Handle(c.Request.Context(), query.FetchNoteByID{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route PUT /api/v1/notes/{id} notes updateNote
// Update a note by ID.
// security:
//   - Bearer:
// responses:
//   200: getNoteResponse
//   400: errorResponse
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	id := c.Param("id")
	var req http_requests.CreateNoteRequest // Reusing CreateNoteRequest as fields are same
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

	n := &note.Note{
		ID:       id,
		UserID:   userID,
		ClientID: req.ClientID,
		Content:  req.Content,
	}

	res, err := h.deps.App.Commands.UpdateNote.Handle(c.Request.Context(), command.UpdateNoteCommand{
		Note: n,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /api/v1/notes/{id} notes deleteNote
// Delete a note by ID.
// security:
//   - Bearer:
// responses:
//   204:
//   500: errorResponse
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	id := c.Param("id")
	err := h.deps.App.Commands.DeleteNote.Handle(c.Request.Context(), command.DeleteNoteCommand{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
