package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type AttendeeRepoPort interface {
	GetAttendee(attendeeID int64) (domain.Attendee, error)
	GetAttendeeEvents(userID int64, params *domain.OrganiserParams) (domain.AttendeeEventFetch, error)
}

type AttendeeRepo struct {
	db AttendeeRepoPort
}

func NewAttendeeRepo(db AttendeeRepoPort) *AttendeeRepo {
	return &AttendeeRepo{db: db}
}

func (r *AttendeeRepo) GetAttendeeByID(attendeeID int64) (domain.Attendee, error) {
	attendee, err := r.db.GetAttendee(attendeeID)
	return attendee, err
}

func (r *AttendeeRepo) GetAttendeeEvents(userID int64, params *domain.OrganiserParams) (domain.AttendeeEventFetch, error) {
	events, err := r.db.GetAttendeeEvents(userID, params)
	return events, err
}
