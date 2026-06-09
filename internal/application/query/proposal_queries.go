package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/proposal"
)

type FetchProposalByID struct {
	ID string
}

type FetchProposalByIDHandler struct {
	repo proposal.ProposalRepositoryInterface
}

func NewFetchProposalByIDHandler(repo proposal.ProposalRepositoryInterface) *FetchProposalByIDHandler {
	return &FetchProposalByIDHandler{repo: repo}
}

func (h *FetchProposalByIDHandler) Handle(ctx context.Context, q FetchProposalByID) (*proposal.Proposal, error) {
	return h.repo.GetByID(q.ID)
}

type FetchProposalList struct{}

type FetchProposalListHandler struct {
	repo proposal.ProposalRepositoryInterface
}

func NewFetchProposalListHandler(repo proposal.ProposalRepositoryInterface) *FetchProposalListHandler {
	return &FetchProposalListHandler{repo: repo}
}

func (h *FetchProposalListHandler) Handle(ctx context.Context, q FetchProposalList) ([]*proposal.Proposal, error) {
	return h.repo.List()
}
