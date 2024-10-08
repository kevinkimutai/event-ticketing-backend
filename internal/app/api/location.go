package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type LocationRepoPort interface {
	GetLocationByID(int64) (domain.Location, error)
	GetLocations() ([]domain.Location, error)
}

type LocationRepo struct {
	db LocationRepoPort
}

func NewLocationRepo(db LocationRepoPort) *LocationRepo {
	return &LocationRepo{db: db}
}

func (r *LocationRepo) GetLocationByID(locationID int64) (domain.Location, error) {
	location, err := r.db.GetLocationByID(locationID)

	return location, err
}
func (r *LocationRepo) GetLocations() ([]domain.Location, error) {
	locations, err := r.db.GetLocations()

	return locations, err
}
