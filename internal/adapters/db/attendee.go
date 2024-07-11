package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
)

func (db *DBAdapter) GetAttendeeByUserID(attendeeID int64) (queries.GetAttendeeByUserIDRow, error) {
	attendee, err := db.queries.GetAttendeeByUserID(context.Background(), attendeeID)

	return attendee, err

}
