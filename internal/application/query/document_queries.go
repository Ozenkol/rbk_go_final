package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type FetchDocumentByID struct {
	ID string
}

type FetchDocumentByIDResult struct {
	Document    *document.Document `json:"document"`
	DownloadURL string             `json:"download_url"`
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

func (h *FetchDocumentByIDHandler) Handle(ctx context.Context, q FetchDocumentByID) (*FetchDocumentByIDResult, error) {
	doc, err := h.repo.GetByID(q.ID)
	if err != nil {
		return nil, err
	}

	downloadURL, err := h.storageProvider.CreateDownloadURL(ctx, doc.ID)
	if err != nil {
		return nil, err
	}

	return &FetchDocumentByIDResult{
		Document:    doc,
		DownloadURL: downloadURL,
	}, nil
}

type FetchDocumentList struct{}

type DocumentWithURL struct {
	*document.Document
	DownloadURL string `json:"download_url"`
}

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

func (h *FetchDocumentListHandler) Handle(ctx context.Context, q FetchDocumentList) ([]DocumentWithURL, error) {
	docs, err := h.repo.List()
	if err != nil {
		return nil, err
	}

	var results []DocumentWithURL
	for _, doc := range docs {
		downloadURL, _ := h.storageProvider.CreateDownloadURL(ctx, doc.ID)
		results = append(results, DocumentWithURL{
			Document:    doc,
			DownloadURL: downloadURL,
		})
	}

	return results, nil
}
