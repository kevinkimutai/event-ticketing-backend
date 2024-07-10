package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type TicketOrderRepoPort interface {
	CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error)
	GetTicketOrders(params domain.Params) ([]domain.TicketOrder, error)
}

type TicketOrderRepo struct {
	db TicketOrderRepoPort
}

func NewTicketOrderRepo(db TicketOrderRepoPort) *TicketOrderRepo {
	return &TicketOrderRepo{db: db}
}

func (r *TicketOrderRepo) CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error) {
	ticketOrder, err := r.db.CreateTicketOrder(order, userID)

	return ticketOrder, err
}

func (r *TicketOrderRepo) GetTicketOrders(params domain.Params) ([]domain.TicketOrder, error) {
	torders, err := r.db.GetTicketOrders(params)

	return torders, err
}
