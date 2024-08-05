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

func (db *DBAdapter) GetLocations() ([]domain.Location, error) {
	ctx := context.Background()

	locs, err := db.queries.ListLocations(ctx)

	var locations []domain.Location

	for _, v := range locs {
		c := domain.Location{
			LocationID: v.LocationID,
			Name:       v.Name,
		}

		locations = append(locations, c)
	}

	return locations, err
}
