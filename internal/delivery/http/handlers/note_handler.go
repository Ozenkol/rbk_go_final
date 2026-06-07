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

type NoteHandler struct {
	deps *http_deps.Dependencies
	logs *slog.Logger
}

func NewNoteHandler(deps *http_deps.Dependencies, logs *slog.Logger) *NoteHandler {
	return &NoteHandler{deps: deps, logs: logs}
}	


// swagger:route POST /api/v1/notes notes createNote
//
// Create a new note.
//
// Consumes:
// - application/json
// Produces:
// - application/json
func (h *NoteHandler) CreateNote(c *gin.Context) {	
	var req http_requests.CreateNoteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	tokenString := c.GetHeader("Authorization")
	userId, err := h.deps.App.Services.AuthService.GetUserByToken(tokenString)
	cmd := command.CreateNoteCommand{
		UserID: userId,
		ClientID: req.ClientID,
		Content: req.Content,
	}
	createdNote, err := h.deps.App.Commands.CreateNote.Handle(cmd)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusFound, createdNote)
}

// swagger:route GET /api/v1/notes/{id} notes getNote
//
// Get a note by ID.
//
// Produces:
// - application/json
// responses:
// - 200: getNoteResponse
func (h *NoteHandler) GetNote(c *gin.Context) {
	q := query.FetchNoteById{
		ID: c.Param("id"),
	}
	note, err := h.deps.App.Queries.GetNoteById.Handle(q)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusFound, note)
}

// swagger:route PUT /api/v1/notes/{id} notes updateNote
//
// Update a note by ID.
//
// Consumes:
// - application/json
// Produces:
// - application/json
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update note - protected route"})
}

// swagger:route DELETE /api/v1/notes/{id} notes deleteNote
//
// Delete a note by ID.
//
// Produces:
// - application/json
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Delete note - protected route"})
}

