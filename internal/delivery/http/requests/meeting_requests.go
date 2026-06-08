package http_requests

import "github.com/Ozenkol/rbk-go-final/internal/domain/meeting"

// CreateMeetingRequest represents the request body for creating a meeting
// swagger:model CreateMeetingRequest
type CreateMeetingRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID        string             `json:"client_id"`
	// Meeting topic
	// example: Project kickoff
	Topic           string             `json:"topic"`
	// Meeting time slot
	TimeSlot        meeting.TimeSlot   `json:"time_slot"`
	// Meeting provider (e.g. Zoom, Google Meet)
	MeetingProvider meeting.MeetingProvider `json:"meeting_provider"`
	// List of attendees
	Attendees       []meeting.Attendee `json:"attendees"`
}

// UpdateMeetingRequest represents the request body for updating a meeting
// swagger:model UpdateMeetingRequest
type UpdateMeetingRequest struct {
	// Client ID
	// example: 550e8400-e29b-41d4-a716-446655440000
	ClientID        string             `json:"client_id"`
	// Meeting topic
	// example: Project kickoff
	Topic           string             `json:"topic"`
	// Meeting time slot
	TimeSlot        meeting.TimeSlot   `json:"time_slot"`
	// Meeting provider
	MeetingProvider meeting.MeetingProvider `json:"meeting_provider"`
	// List of attendees
	Attendees       []meeting.Attendee `json:"attendees"`
}
