package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type CreateDocumentCommand struct {
	Document *document.Document
}

type CreateDocumentResult struct {
	Document  *document.Document `json:"document"`
	UploadURL string             `json:"upload_url"`
}

type CreateDocumentHandler struct {
	repo            document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewCreateDocumentHandler(
	repo document.DocumentRepositoryInterface,
	storageProvider adapters.ObjectStorageProvider,
) *CreateDocumentHandler {
	return &CreateDocumentHandler{
		repo:            repo,
		storageProvider: storageProvider,
	}
}

func (h *CreateDocumentHandler) Handle(ctx context.Context, cmd CreateDocumentCommand) (*CreateDocumentResult, error) {
	doc, err := h.repo.Create(cmd.Document)
	if err != nil {
		return nil, err
	}

	uploadURL, err := h.storageProvider.CreateUploadURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	downloadURL, err := h.storageProvider.CreateDownloadURL(ctx, doc.ID)
	if err == nil {
		doc.StorageReference.URL = downloadURL
		doc, err = h.repo.Update(doc)
		if err != nil {
			return nil, err
		}
	}

	return &CreateDocumentResult{
		Document:  doc,
		UploadURL: uploadURL,
	}, nil
}

type UpdateDocumentCommand struct {
	Document *document.Document
}

type UpdateDocumentResult struct {
	Document  *document.Document `json:"document"`
	UpdateURL string             `json:"update_url"`
}

type UpdateDocumentHandler struct {
	repo            document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewUpdateDocumentHandler(
	repo document.DocumentRepositoryInterface,
	storageProvider adapters.ObjectStorageProvider,
) *UpdateDocumentHandler {
	return &UpdateDocumentHandler{
		repo:            repo,
		storageProvider: storageProvider,
	}
}

func (h *UpdateDocumentHandler) Handle(ctx context.Context, cmd UpdateDocumentCommand) (*UpdateDocumentResult, error) {
	doc, err := h.repo.Update(cmd.Document)
	if err != nil {
		return nil, err
	}

	updateURL, err := h.storageProvider.CreateUpdateURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	return &UpdateDocumentResult{
		Document:  doc,
		UpdateURL: updateURL,
	}, nil
}

type DeleteDocumentCommand struct {
	ID string
}

type DeleteDocumentHandler struct {
	repo            document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewDeleteDocumentHandler(
	repo document.DocumentRepositoryInterface,
	storageProvider adapters.ObjectStorageProvider,
) *DeleteDocumentHandler {
	return &DeleteDocumentHandler{
		repo:            repo,
		storageProvider: storageProvider,
	}
}

func (h *DeleteDocumentHandler) Handle(ctx context.Context, cmd DeleteDocumentCommand) error {
	if err := h.storageProvider.DeleteObject(ctx, cmd.ID); err != nil {
		return err
	}
	return h.repo.Delete(cmd.ID)
}
