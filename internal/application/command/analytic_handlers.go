package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/analytic"
)

type CreateAnalyticCommand struct {
	Analytic *analytic.Analytic
}

type CreateAnalyticHandler struct {
	repo analytic.AnalyticRepositoryInterface
}

func NewCreateAnalyticHandler(repo analytic.AnalyticRepositoryInterface) *CreateAnalyticHandler {
	return &CreateAnalyticHandler{repo: repo}
}

func (h *CreateAnalyticHandler) Handle(ctx context.Context, cmd CreateAnalyticCommand) (*analytic.Analytic, error) {
	return h.repo.Create(cmd.Analytic)
}

type UpdateAnalyticCommand struct {
	Analytic *analytic.Analytic
}

type UpdateAnalyticHandler struct {
	repo analytic.AnalyticRepositoryInterface
}

func NewUpdateAnalyticHandler(repo analytic.AnalyticRepositoryInterface) *UpdateAnalyticHandler {
	return &UpdateAnalyticHandler{repo: repo}
}

func (h *UpdateAnalyticHandler) Handle(ctx context.Context, cmd UpdateAnalyticCommand) (*analytic.Analytic, error) {
	return h.repo.Update(cmd.Analytic)
}

type DeleteAnalyticCommand struct {
	ID string
}

type DeleteAnalyticHandler struct {
	repo analytic.AnalyticRepositoryInterface
}

func NewDeleteAnalyticHandler(repo analytic.AnalyticRepositoryInterface) *DeleteAnalyticHandler {
	return &DeleteAnalyticHandler{repo: repo}
}

func (h *DeleteAnalyticHandler) Handle(ctx context.Context, cmd DeleteAnalyticCommand) error {
	return h.repo.Delete(cmd.ID)
}
