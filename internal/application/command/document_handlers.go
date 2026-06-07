package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type CreateDocumentCommand struct {
	Document *document.Document
}

type CreateDocumentHandler struct {
	repo document.DocumentRepositoryInterface
}

func NewCreateDocumentHandler(repo document.DocumentRepositoryInterface) *CreateDocumentHandler {
	return &CreateDocumentHandler{repo: repo}
}

func (h *CreateDocumentHandler) Handle(ctx context.Context, cmd CreateDocumentCommand) (*document.Document, error) {
	return h.repo.Create(cmd.Document)
}

type UpdateDocumentCommand struct {
	Document *document.Document
}

type UpdateDocumentHandler struct {
	repo document.DocumentRepositoryInterface
}

func NewUpdateDocumentHandler(repo document.DocumentRepositoryInterface) *UpdateDocumentHandler {
	return &UpdateDocumentHandler{repo: repo}
}

func (h *UpdateDocumentHandler) Handle(ctx context.Context, cmd UpdateDocumentCommand) (*document.Document, error) {
	return h.repo.Update(cmd.Document)
}

type DeleteDocumentCommand struct {
	ID string
}

type DeleteDocumentHandler struct {
	repo document.DocumentRepositoryInterface
}

func NewDeleteDocumentHandler(repo document.DocumentRepositoryInterface) *DeleteDocumentHandler {
	return &DeleteDocumentHandler{repo: repo}
}

func (h *DeleteDocumentHandler) Handle(ctx context.Context, cmd DeleteDocumentCommand) error {
	return h.repo.Delete(cmd.ID)
}
