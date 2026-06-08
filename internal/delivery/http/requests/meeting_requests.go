package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/meeting"

type CreateMeetingRequest struct {
	ClientID        string             `json:"client_id"`
	CompanyID       string             `json:"company_id"`
	Topic           string             `json:"topic"`
	TimeSlot        meeting.TimeSlot   `json:"time_slot"`
	MeetingProvider meeting.MeetingProvider `json:"meeting_provider"`
	Attendees       []meeting.Attendee `json:"attendees"`
}

type UpdateMeetingRequest struct {
	ClientID        string             `json:"client_id"`
	CompanyID       string             `json:"company_id"`
	Topic           string             `json:"topic"`
	TimeSlot        meeting.TimeSlot   `json:"time_slot"`
	MeetingProvider meeting.MeetingProvider `json:"meeting_provider"`
	Attendees       []meeting.Attendee `json:"attendees"`
}
