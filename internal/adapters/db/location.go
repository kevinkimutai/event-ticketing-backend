package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) GetLocationByID(locationID int64) (domain.Location, error) {
	ctx := context.Background()

	loc, err := db.queries.GetLocation(ctx, locationID)

	return domain.Location{
		LocationID: loc.LocationID,
		Name:       loc.Name,
	}, err
}
