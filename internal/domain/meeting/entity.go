package meeting

type Meeting struct {
	ID       string
	ClientID string
	CompanyID string
	Topic    string
	TimeSlot TimeSlot
	MeetingProvider MeetingProvider
	Attendees []Attendee
}
