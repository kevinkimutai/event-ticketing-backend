package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) CreateEvent(event *domain.Event) (domain.Event, error) {
	ctx := context.Background()
	e, err := db.queries.CreateEvent(ctx, queries.CreateEventParams{
		Name:        event.Name,
		CategoryID:  event.CategoryID,
		Date:        utils.ConvertTimeToTimestamp(event.Date),
		FromTime:    utils.ConvertTimeToTimestamp(event.FromTime),
		ToTime:      utils.ConvertTimeToTimestamp(event.ToTime),
		Location:    event.Location,
		Description: utils.ConvertStringToText(event.Description),
		Longitude:   event.Longitude,
		Latitude:    event.Latitude,
	})

	return domain.Event{
		EventID:     e.EventID,
		Name:        e.Name,
		CategoryID:  e.CategoryID,
		Date:        e.Date.Time,
		FromTime:    e.FromTime.Time,
		ToTime:      e.ToTime.Time,
		Location:    e.Location,
		Description: e.Description.String,
		Longitude:   e.Longitude,
		Latitude:    e.Latitude,
	}, err
}
