package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
	"github.com/kevinkimutai/ticketingapp/internal/utils"
)

func (db *DBAdapter) CreateEvent(event *domain.Event, userID int64) (domain.Event, error) {
	ctx := context.Background()

	//Start TX
	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.Event{}, errors.New("failed to start tx")
	}

	qtx := db.queries.WithTx(tx)

	//Create Event
	e, err := qtx.CreateEvent(ctx, queries.CreateEventParams{
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
	if err != nil {
		tx.Rollback(ctx)
		return domain.Event{}, err
	}

	//Create Organiser
	_, err = qtx.CreateOrganiser(ctx, queries.CreateOrganiserParams{
		UserID:  userID,
		EventID: e.EventID,
	})
	if err != nil {
		tx.Rollback(ctx)
		return domain.Event{}, err
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return domain.Event{}, err
	}

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
