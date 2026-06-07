package meeting

import "errors"

type TimeSlot struct {
	StartTime string
	EndTime   string
	TimeZone string
}

type Attendee struct {
	UserID string
	Email string
}

type MeetingProvider struct{
	ProviderName string
	URL string
}

func (*TimeSlot) NewTimeSlot(startTime, endTime, timeZone string) (*TimeSlot, error) {
	if startTime == "" || endTime == "" || timeZone == "" {
		return nil, errors.New("start time, end time, and time zone must be provided")
	}
	return &TimeSlot{
		StartTime: startTime,
		EndTime:   endTime,
		TimeZone: timeZone,
	}, nil
}

func (*Attendee) NewAttendee(userID, email string) (*Attendee, error) {
	if userID == "" || email == "" {
		return nil, errors.New("user ID and email must be provided")
	}
	return &Attendee{
		UserID: userID,
		Email: email,
	}, nil
}

func (*MeetingProvider) NewMeetingProvider(providerName, url string) (*MeetingProvider, error) {
	if providerName == "" || url == "" {
		return nil, errors.New("provider name and URL must be provided")
	}
	return &MeetingProvider{
		ProviderName: providerName,
		URL: url,
	}, nil
}