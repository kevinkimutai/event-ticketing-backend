package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type TicketTypeRepoPort interface {
	CreateTicketType(*domain.TicketType) (domain.TicketType, error)
	GetTicketTypesByEvent(eventID int64) ([]domain.TicketType, error)
}

type TicketTypeRepo struct {
	db TicketTypeRepoPort
}

func NewTicketTypeRepo(db TicketTypeRepoPort) *TicketTypeRepo {
	return &TicketTypeRepo{db: db}
}

func (r *TicketTypeRepo) CreateTicketType(t *domain.TicketType) (domain.TicketType, error) {
	ttype, err := r.db.CreateTicketType(t)

	return ttype, err
}

func (r *TicketTypeRepo) GetTicketTypesByEvent(eventID int64) ([]domain.TicketType, error) {
	ticketTypes, err := r.db.GetTicketTypesByEvent(eventID)

	return ticketTypes, err
}
