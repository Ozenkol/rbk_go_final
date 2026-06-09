package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/proposal"
)

type CreateProposalCommand struct {
	Proposal *proposal.Proposal
}

type CreateProposalHandler struct {
	repo proposal.ProposalRepositoryInterface
}

func NewCreateProposalHandler(repo proposal.ProposalRepositoryInterface) *CreateProposalHandler {
	return &CreateProposalHandler{repo: repo}
}

func (h *CreateProposalHandler) Handle(ctx context.Context, cmd CreateProposalCommand) (*proposal.Proposal, error) {
	return h.repo.Create(cmd.Proposal)
}

type UpdateProposalCommand struct {
	Proposal *proposal.Proposal
}

type UpdateProposalHandler struct {
	repo proposal.ProposalRepositoryInterface
}

func NewUpdateProposalHandler(repo proposal.ProposalRepositoryInterface) *UpdateProposalHandler {
	return &UpdateProposalHandler{repo: repo}
}

func (h *UpdateProposalHandler) Handle(ctx context.Context, cmd UpdateProposalCommand) (*proposal.Proposal, error) {
	return h.repo.Update(cmd.Proposal)
}

type DeleteProposalCommand struct {
	ID string
}

type DeleteProposalHandler struct {
	repo proposal.ProposalRepositoryInterface
}

func NewDeleteProposalHandler(repo proposal.ProposalRepositoryInterface) *DeleteProposalHandler {
	return &DeleteProposalHandler{repo: repo}
}

func (h *DeleteProposalHandler) Handle(ctx context.Context, cmd DeleteProposalCommand) error {
	return h.repo.Delete(cmd.ID)
}
