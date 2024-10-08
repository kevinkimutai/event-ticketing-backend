package db

import (
	"context"
	"math"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
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

func (db *DBAdapter) GetAttendeeEvents(userID int64, params *domain.OrganiserParams) (domain.AttendeeEventFetch, error) {
	ctx := context.Background()

	events, err := db.queries.GetAttendeeEvents(ctx, queries.GetAttendeeEventsParams{
		UserID: userID,
		Limit:  params.Limit,
		Offset: params.Page,
	})
	if err != nil {
		return domain.AttendeeEventFetch{}, err
	}

	eAttended, err := db.queries.GetEventsAttended(ctx, userID)
	if err != nil {
		return domain.AttendeeEventFetch{}, err
	}

	count, err := db.queries.GetCountAttendeeEvents(ctx, userID)
	if err != nil {
		return domain.AttendeeEventFetch{}, err
	}

	//Get Page
	page := getPage(params.Page, params.Limit)

	var attendeeEvents []domain.AttendeeEvents

	for _, v := range events {
		e := domain.AttendeeEvents{
			AttendeeID:  v.AttendeeID,
			EventName:   v.EventName,
			EventDate:   v.EventDate.Time,
			PaymentID:   v.PaymentID,
			Quantity:    v.Quantity,
			TotalAmount: utils.ConvertNumericToFloat64(v.TotalAmount),
		}

		attendeeEvents = append(attendeeEvents, e)

	}

	//TODO:will add pagination
	return domain.AttendeeEventFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64((params.Limit)))),
		Total:         count,
		Data: domain.AttendeeData{
			TotalEvents: eAttended.EventsAttended,
			TotalSpent:  eAttended.TotalSpent,
			Data:        attendeeEvents,
		},
	}, nil

}
