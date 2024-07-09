package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type EventRepoPort interface {
	CreateEvent(event *domain.Event, userID int64) (domain.Event, error)
	GetEvents(domain.Params) (domain.EventsFetch, error)
}

type EventRepo struct {
	db EventRepoPort
}

func NewEventRepo(db EventRepoPort) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r *EventRepo) CreateEvent(event *domain.Event, userID int64) (domain.Event, error) {
	e, err := r.db.CreateEvent(event, userID)

	return e, err
}

func (r *EventRepo) GetEvents(params domain.Params) (domain.EventsFetch, error) {
	events, err := r.db.GetEvents(params)

	return events, err
}
