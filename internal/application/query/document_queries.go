package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type FetchDocumentByID struct {
	ID string
}

type FetchDocumentByIDHandler struct {
	repo            document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewFetchDocumentByIDHandler(
	repo document.DocumentRepositoryInterface,
	storageProvider adapters.ObjectStorageProvider,
) *FetchDocumentByIDHandler {
	return &FetchDocumentByIDHandler{
		repo:            repo,
		storageProvider: storageProvider,
	}
}

func (h *FetchDocumentByIDHandler) Handle(ctx context.Context, q FetchDocumentByID) (*application_shared.DocumentResponse, error) {
	doc, err := h.repo.GetByID(q.ID)
	if err != nil {
		return nil, err
	}

	uploadURL, err := h.storageProvider.CreateUploadURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	downloadURL, err := h.storageProvider.CreateDownloadURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	updateURL, err := h.storageProvider.CreateUpdateURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	return &application_shared.DocumentResponse{
		Document:    doc,
		UploadURL:   uploadURL,
		DownloadURL: downloadURL,
		UpdateURL:   updateURL,
	}, nil
}

type FetchDocumentList struct{}

type FetchDocumentListHandler struct {
	repo            document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewFetchDocumentListHandler(
	repo document.DocumentRepositoryInterface,
	storageProvider adapters.ObjectStorageProvider,
) *FetchDocumentListHandler {
	return &FetchDocumentListHandler{
		repo:            repo,
		storageProvider: storageProvider,
	}
}

func (h *FetchDocumentListHandler) Handle(ctx context.Context, q FetchDocumentList) ([]application_shared.DocumentResponse, error) {
	docs, err := h.repo.List()
	if err != nil {
		return nil, err
	}

	var results []application_shared.DocumentResponse
	for _, doc := range docs {
		uploadURL, _ := h.storageProvider.CreateUploadURL(ctx, doc.ID)
		downloadURL, _ := h.storageProvider.CreateDownloadURL(ctx, doc.ID)
		updateURL, _ := h.storageProvider.CreateUpdateURL(ctx, doc.ID)

		results = append(results, application_shared.DocumentResponse{
			Document:    doc,
			UploadURL:   uploadURL,
			DownloadURL: downloadURL,
			UpdateURL:   updateURL,
		})
	}

	return results, nil
}
