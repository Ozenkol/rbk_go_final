package service

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type DocumentService struct {
	documentRepo    document.DocumentRepositoryInterface
	storageProvider adapters.ObjectStorageProvider
}

func NewDocumentService(repo document.DocumentRepositoryInterface, storage adapters.ObjectStorageProvider) *DocumentService {
	return &DocumentService{
		documentRepo:    repo,
		storageProvider: storage,
	}
}

func (s *DocumentService) CreateUploadURL(ctx context.Context, key string) (string, error) {
	return s.storageProvider.CreateUploadURL(ctx, key)
}

func (s *DocumentService) CreateDownloadURL(ctx context.Context, key string) (string, error) {
	return s.storageProvider.CreateDownloadURL(ctx, key)
}

func (s *DocumentService) SaveDocument(doc *document.Document) (*document.Document, error) {
	return s.documentRepo.Create(doc)
}

func (s *DocumentService) UpdateDocument(doc *document.Document) (*document.Document, error) {
	return s.documentRepo.Update(doc)
}
