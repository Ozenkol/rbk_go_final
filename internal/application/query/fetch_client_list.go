package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/client"
)

type ClientFilter struct {
	Name        string
	Email       string
	PhoneNumber string
}

type FetchClientListQuery struct {
	Filter     ClientFilter
	Pagination application_shared.Pagination
}

type ClientQueryRepositoryInterface interface {
	FindByFilter(filter ClientFilter, pagination application_shared.Pagination) ([]*client.Client, error)
}

type FetchClientListHandler struct {
	clientQueryRepository ClientQueryRepositoryInterface
}

func NewFetchClientListHandler(clientRepo ClientQueryRepositoryInterface) *FetchClientListHandler {
	return &FetchClientListHandler{clientQueryRepository: clientRepo}
}

func (h *FetchClientListHandler) Handle(query FetchClientListQuery) ([]*client.Client, error) {
	clients, err := h.clientQueryRepository.FindByFilter(query.Filter, query.Pagination)
	if err != nil {
		return nil, err
	}
	return clients, nil	
}