package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type EventRepoPort interface {
	CreateEvent(*domain.Event) (domain.Event, error)
}

type EventRepo struct {
	db EventRepoPort
}

func NewEventRepo(db EventRepoPort) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r *EventRepo) CreateEvent(event *domain.Event) (domain.Event, error) {
	e, err := r.db.CreateEvent(event)

	return e, err
}
