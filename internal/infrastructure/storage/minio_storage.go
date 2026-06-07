package storage

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/application/adapters"
)

type MinioStorage struct {
	// In a real implementation, we would have minio.Client here
}

func NewMinioStorage() adapters.ObjectStorageProvider {
	return &MinioStorage{}
}

func (s *MinioStorage) CreateUploadURL(ctx context.Context, key string) (string, error) {
	return "http://localhost:9000/upload/" + key, nil
}

func (s *MinioStorage) CreateDownloadURL(ctx context.Context, key string) (string, error) {
	return "http://localhost:9000/download/" + key, nil
}

func (s *MinioStorage) DeleteObject(ctx context.Context, key string) error {
	return nil
}
