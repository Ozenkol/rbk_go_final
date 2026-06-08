package http_requests

type CreateFileRequest struct {
	ClientID         string `json:"client_id"`
	CompanyID        string `json:"company_id"`
	StorageReference string `json:"storage_reference"`
	Name             string `json:"name"`
}

type UpdateFileRequest struct {
	ClientID         string `json:"client_id"`
	CompanyID        string `json:"company_id"`
	StorageReference string `json:"storage_reference"`
	Name             string `json:"name"`
}
