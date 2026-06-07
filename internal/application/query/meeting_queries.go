package query

import (
	"context"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
)

type FetchMeetingByID struct {
	ID string
}

type FetchMeetingByIDHandler struct {
	repo meeting.MeetingRepositoryInterface
}

func NewFetchMeetingByIDHandler(repo meeting.MeetingRepositoryInterface) *FetchMeetingByIDHandler {
	return &FetchMeetingByIDHandler{repo: repo}
}

func (h *FetchMeetingByIDHandler) Handle(ctx context.Context, q FetchMeetingByID) (*meeting.Meeting, error) {
	return h.repo.GetByID(q.ID)
}

type FetchMeetingList struct{}

type FetchMeetingListHandler struct {
	repo meeting.MeetingRepositoryInterface
}

func NewFetchMeetingListHandler(repo meeting.MeetingRepositoryInterface) *FetchMeetingListHandler {
	return &FetchMeetingListHandler{repo: repo}
}

func (h *FetchMeetingListHandler) Handle(ctx context.Context, q FetchMeetingList) ([]*meeting.Meeting, error) {
	return h.repo.List()
}
