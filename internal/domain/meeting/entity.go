package meeting

type Meeting struct {
	ID        string
	UserID    string
	ClientID  string
	CompanyID string
	Topic     string
	TimeSlot TimeSlot
	MeetingProvider MeetingProvider
	Attendees []Attendee
}
