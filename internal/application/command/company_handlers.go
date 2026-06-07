package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/company"
)

type CreateCompanyCommand struct {
	Company *company.Company
}

type CreateCompanyHandler struct {
	repo company.CompanyRepositoryInterface
}

func NewCreateCompanyHandler(repo company.CompanyRepositoryInterface) *CreateCompanyHandler {
	return &CreateCompanyHandler{repo: repo}
}

func (h *CreateCompanyHandler) Handle(ctx context.Context, cmd CreateCompanyCommand) (*company.Company, error) {
	return h.repo.Create(cmd.Company)
}

type UpdateCompanyCommand struct {
	Company *company.Company
}

type UpdateCompanyHandler struct {
	repo company.CompanyRepositoryInterface
}

func NewUpdateCompanyHandler(repo company.CompanyRepositoryInterface) *UpdateCompanyHandler {
	return &UpdateCompanyHandler{repo: repo}
}

func (h *UpdateCompanyHandler) Handle(ctx context.Context, cmd UpdateCompanyCommand) (*company.Company, error) {
	return h.repo.Update(cmd.Company)
}

type DeleteCompanyCommand struct {
	ID string
}

type DeleteCompanyHandler struct {
	repo company.CompanyRepositoryInterface
}

func NewDeleteCompanyHandler(repo company.CompanyRepositoryInterface) *DeleteCompanyHandler {
	return &DeleteCompanyHandler{repo: repo}
}

func (h *DeleteCompanyHandler) Handle(ctx context.Context, cmd DeleteCompanyCommand) error {
	return h.repo.Delete(cmd.ID)
}
