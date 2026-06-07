package adapters

import "context"

type UploadURLProvider interface {
	CreateUploadURL(
		ctx context.Context,
		key string,
	) (string, error)
}

type DownloadURLProvider interface {
	CreateDownloadURL(
		ctx context.Context,
		key string,
	) (string, error)
}

type ObjectDeleter interface {
	DeleteObject(
		ctx context.Context,
		key string,
	) error
}

type ObjectStorageProvider interface {
	UploadURLProvider
	DownloadURLProvider
	ObjectDeleter
}