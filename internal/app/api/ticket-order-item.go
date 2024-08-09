package api

import "github.com/kevinkimutai/ticketingapp/internal/app/domain"

type TicketOrderItemPort interface {
	GetTicketOrder(orderItemId int64) (domain.TicketOrder, error)
}

type TicketOrderItemRepo struct {
	db TicketOrderItemPort
}

func NewTicketOrderItemRepo(db TicketOrderItemPort) *TicketOrderItemRepo {
	return &TicketOrderItemRepo{db: db}
}
