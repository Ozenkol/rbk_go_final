package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/contract"
)

type CreateContractCommand struct {
	Contract *contract.Contract
}

type CreateContractHandler struct {
	repo contract.ContractRepositoryInterface
}

func NewCreateContractHandler(repo contract.ContractRepositoryInterface) *CreateContractHandler {
	return &CreateContractHandler{repo: repo}
}

func (h *CreateContractHandler) Handle(ctx context.Context, cmd CreateContractCommand) (*contract.Contract, error) {
	return h.repo.Create(cmd.Contract)
}

type UpdateContractCommand struct {
	Contract *contract.Contract
}

type UpdateContractHandler struct {
	repo contract.ContractRepositoryInterface
}

func NewUpdateContractHandler(repo contract.ContractRepositoryInterface) *UpdateContractHandler {
	return &UpdateContractHandler{repo: repo}
}

func (h *UpdateContractHandler) Handle(ctx context.Context, cmd UpdateContractCommand) (*contract.Contract, error) {
	return h.repo.Update(cmd.Contract)
}

type DeleteContractCommand struct {
	ID string
}

type DeleteContractHandler struct {
	repo contract.ContractRepositoryInterface
}

func NewDeleteContractHandler(repo contract.ContractRepositoryInterface) *DeleteContractHandler {
	return &DeleteContractHandler{repo: repo}
}

func (h *DeleteContractHandler) Handle(ctx context.Context, cmd DeleteContractCommand) error {
	return h.repo.Delete(cmd.ID)
}
