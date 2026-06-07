package command

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
)

type CreateMeetingCommand struct {
	Meeting *meeting.Meeting
}

type CreateMeetingHandler struct {
	repo meeting.MeetingRepositoryInterface
}

func NewCreateMeetingHandler(repo meeting.MeetingRepositoryInterface) *CreateMeetingHandler {
	return &CreateMeetingHandler{repo: repo}
}

func (h *CreateMeetingHandler) Handle(ctx context.Context, cmd CreateMeetingCommand) (*meeting.Meeting, error) {
	return h.repo.Create(cmd.Meeting)
}

type UpdateMeetingCommand struct {
	Meeting *meeting.Meeting
}

type UpdateMeetingHandler struct {
	repo meeting.MeetingRepositoryInterface
}

func NewUpdateMeetingHandler(repo meeting.MeetingRepositoryInterface) *UpdateMeetingHandler {
	return &UpdateMeetingHandler{repo: repo}
}

func (h *UpdateMeetingHandler) Handle(ctx context.Context, cmd UpdateMeetingCommand) (*meeting.Meeting, error) {
	return h.repo.Update(cmd.Meeting)
}

type DeleteMeetingCommand struct {
	ID string
}

type DeleteMeetingHandler struct {
	repo meeting.MeetingRepositoryInterface
}

func NewDeleteMeetingHandler(repo meeting.MeetingRepositoryInterface) *DeleteMeetingHandler {
	return &DeleteMeetingHandler{repo: repo}
}

func (h *DeleteMeetingHandler) Handle(ctx context.Context, cmd DeleteMeetingCommand) error {
	return h.repo.Delete(cmd.ID)
}
