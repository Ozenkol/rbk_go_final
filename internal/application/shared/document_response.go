package application_shared

import "github.com/Ozenkol/rbk-go-final/internal/domain/document"

// swagger:model DocumentResponse
type DocumentResponse struct {
	*document.Document
	UploadURL   string `json:"upload_url,omitempty"`
	DownloadURL string `json:"download_url,omitempty"`
	UpdateURL   string `json:"update_url,omitempty"`
}
