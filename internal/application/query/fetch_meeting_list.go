package query

import (
	application_shared "github.com/Ozenkol/rbk-go-final/internal/application/shared"
	"github.com/Ozenkol/rbk-go-final/internal/domain/meeting"
)

type MeetingFilter struct {
	ClientID string
	Topic    string
	TimeSlot string
}

type MeetingQueryRepositoryInterface interface {
	GetByID(id string) (*meeting.Meeting, error)
	FindByFilter(filter MeetingFilter, pagination application_shared.Pagination) ([]*meeting.Meeting, error)
}

type FetchMeetingListQuery struct {
	Filter     MeetingFilter
	Pagination application_shared.Pagination
}

type FetchMeetingListHandler struct {
	meetingQueryRepository MeetingQueryRepositoryInterface
}

func NewFetchMeetingListHandler(meetingRepo MeetingQueryRepositoryInterface) *FetchMeetingListHandler {	
	return &FetchMeetingListHandler{meetingQueryRepository: meetingRepo}
}

func (h *FetchMeetingListHandler) Handle(query FetchMeetingListQuery) ([]*meeting.Meeting, error) {
	meetings, err := h.meetingQueryRepository.FindByFilter(query.Filter, query.Pagination)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}