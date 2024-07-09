package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type EventRepoPort interface {
	CreateEvent(event *domain.Event, userID int64) (domain.Event, error)
	GetEvents(domain.Params) (domain.EventsFetch, error)
	GetEventByID(int64) (domain.Event, error)
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

func (r *EventRepo) GetEventByID(eventID int64) (domain.Event, error) {
	event, err := r.db.GetEventByID(eventID)

	return event, err
}
