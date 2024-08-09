package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) GetAttendeeByUserID(attendeeID int64) (queries.GetAttendeeByUserIDRow, error) {
	attendee, err := db.queries.GetAttendeeByUserID(context.Background(), attendeeID)

	return attendee, err

}

func (db *DBAdapter) GetAttendee(attendeeID int64) (domain.Attendee, error) {
	attendee, err := db.queries.GetAttendee(context.Background(), attendeeID)

	return domain.Attendee{
		AttendeeID: attendee.AttendeeID,
		UserID:     attendee.UserID,
	}, err

}
