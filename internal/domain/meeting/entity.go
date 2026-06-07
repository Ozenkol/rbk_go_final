package meeting

type Meeting struct {
	ID       string
	ClientID string
	Topic    string
	TimeSlot TimeSlot
	MeetingProvider MeetingProvider
	Attendees []Attendee
}
