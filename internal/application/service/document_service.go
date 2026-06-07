package service

import (
	"context"

	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
	"github.com/Ozenkol/rbk-go-final/internal/domain/document"
)

type DocuemntService struct {
	objectStorageProvider adapters.ObjectStorageProvider
	documentRepo document.DocumentRepositoryInterface
}

func NewDocumentService(storageProvider adapters.ObjectStorageProvider, documentRepo document.DocumentRepositoryInterface) *DocuemntService {
	return &DocuemntService{
		objectStorageProvider: storageProvider,
		documentRepo: documentRepo,
	}
}

func (s *DocuemntService) UploadDocument(ctx context.Context, document *document.Document) (string, error) {
	createdDocument, err := s.documentRepo.Create(document)
	if err != nil {
		return "", err
	}
	storageRef, err := s.objectStorageProvider.CreateUploadURL(ctx, createdDocument.ID)
	if err != nil {
		return "", err
	}
	return storageRef, nil
}

func (s *DocuemntService) GetDocument(ctx context.Context, documentID string) (string, error) {
	doc, err := s.documentRepo.GetByID(documentID)
	if err != nil {
		return "", err
	}
	downloadURL, err := s.objectStorageProvider.CreateDownloadURL(ctx, doc.ID)
	if err != nil {
		return "", err
	}
	return downloadURL, nil
}

func (s *DocuemntService) UpdateDocument(ctx context.Context, document *document.Document) (string, error) {
	err := s.documentRepo.Update(document)
	if err != nil {
		return "", err
	}
	downloadURL, err := s.objectStorageProvider.CreateDownloadURL(ctx, document.ID)
	if err != nil {
		return "", err
	}
	return downloadURL, nil
}

func (s *DocuemntService) DeleteDocument(ctx context.Context, documentID string) error {
	existingDoc, err := s.documentRepo.GetByID(documentID)
	if err != nil {
		return err
	}
	err = s.objectStorageProvider.DeleteObject(ctx, existingDoc.ID)
	if err != nil {
		return err
	}
	return s.documentRepo.Delete(documentID)
}