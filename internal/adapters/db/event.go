package db

import (
	"context"
	"database/sql"
	"errors"
	"math"

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
		PosterUrl:   event.PosterUrl,
		LocationID:  event.LocationID,
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
		PosterUrl:   e.PosterUrl,
	}, nil
}

func (db *DBAdapter) GetEvents(params *domain.Params) (domain.EventsFetch, error) {
	ctx := context.Background()

	var eParams = queries.ListUpcomingEventsParams{
		Limit:  params.Limit,
		Offset: params.Page,
	}

	if params.CategoryID != 0 {
		eParams.CategoryID = sql.NullInt64{Int64: params.CategoryID, Valid: true}
	}

	if params.LocationID != 0 {
		eParams.LocationID = sql.NullInt64{Int64: params.LocationID, Valid: true}
	}

	// fmt.Println(eParams.CategoryID, eParams.LocationID)
	//Get Products
	events, err := db.queries.ListUpcomingEvents(ctx, eParams)
	if err != nil {
		return domain.EventsFetch{}, err

	}

	//Get Count
	count, err := db.queries.GetTotalEventsCount(ctx)
	if err != nil {
		return domain.EventsFetch{}, err

	}

	//Get Page
	page := getPage(params.Page, params.Limit)

	//map struct
	var evs []domain.Event

	for _, item := range events {

		event := domain.Event{
			EventID:     item.EventID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description.String,
			PosterUrl:   item.PosterUrl,
			Location:    item.Location,
			Longitude:   item.Longitude,
			Latitude:    item.Latitude,
			Date:        item.Date.Time,
			FromTime:    item.FromTime.Time,
			ToTime:      item.ToTime.Time,
			CreatedAt:   item.CreatedAt.Time,
		}
		// Append the struct to the struct array
		evs = append(evs, event)
	}

	return domain.EventsFetch{
		Page:          page,
		NumberOfPages: uint(math.Ceil(float64(count) / float64((params.Limit)))),
		Total:         uint(count),
		Data:          evs,
	}, nil

}

func getPage(offset, limit int32) uint {
	return uint((offset / limit) + 1)
}

func (db *DBAdapter) GetEventByID(eventID int64) (domain.Event, error) {
	ctx := context.Background()

	e, err := db.queries.GetEvent(ctx, eventID)
	if err != nil {
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
		PosterUrl:   e.PosterUrl,
		LocationID:  e.LocationID,
	}, nil
}
