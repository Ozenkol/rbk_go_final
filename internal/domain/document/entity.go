package document

import "github.com/Ozenkol/rbk-go-final/internal/domain/shared"

type Document struct {
	ID               string
	ClientID         string
	Type             string
	Number           string
	IssuedBy         string
	StorageReference shared.StorageReference
	IssuedDate       string
	ExpirationDate   string
}
