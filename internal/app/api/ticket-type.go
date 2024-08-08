package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type TicketTypeRepoPort interface {
	CreateTicketType(t *domain.TicketType, eventID int64) (domain.TicketType, error)
	GetTicketTypesByEvent(eventID int64) ([]domain.TicketType, error)
	GetTicket(ticketTypeID int64) (domain.Ticket, error)
}

type TicketTypeRepo struct {
	db TicketTypeRepoPort
}

func NewTicketTypeRepo(db TicketTypeRepoPort) *TicketTypeRepo {
	return &TicketTypeRepo{db: db}
}

func (r *TicketTypeRepo) CreateTicketType(t *domain.TicketType, eventID int64) (domain.TicketType, error) {
	ttype, err := r.db.CreateTicketType(t, eventID)

	return ttype, err
}

func (r *TicketTypeRepo) GetTicketTypesByEvent(eventID int64) ([]domain.TicketType, error) {
	ticketTypes, err := r.db.GetTicketTypesByEvent(eventID)

	return ticketTypes, err
}

func (r *TicketTypeRepo) GetTicket(ticketTypeID int64) (domain.Ticket, error) {
	ticket, err := r.db.GetTicket(ticketTypeID)

	return ticket, err
}
