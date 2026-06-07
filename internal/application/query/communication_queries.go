package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/communication"
)

type FetchCommunicationByID struct {
	ID string
}

type FetchCommunicationByIDHandler struct {
	repo communication.CommunicationRepositoryInterface
}

func NewFetchCommunicationByIDHandler(repo communication.CommunicationRepositoryInterface) *FetchCommunicationByIDHandler {
	return &FetchCommunicationByIDHandler{repo: repo}
}

func (h *FetchCommunicationByIDHandler) Handle(ctx context.Context, q FetchCommunicationByID) (*communication.Communication, error) {
	return h.repo.GetByID(q.ID)
}

type FetchCommunicationList struct{}

type FetchCommunicationListHandler struct {
	repo communication.CommunicationRepositoryInterface
}

func NewFetchCommunicationListHandler(repo communication.CommunicationRepositoryInterface) *FetchCommunicationListHandler {
	return &FetchCommunicationListHandler{repo: repo}
}

func (h *FetchCommunicationListHandler) Handle(ctx context.Context, q FetchCommunicationList) ([]*communication.Communication, error) {
	return h.repo.List()
}
