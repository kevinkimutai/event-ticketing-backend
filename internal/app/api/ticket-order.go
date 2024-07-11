package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketOrderRepoPort interface {
	CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error)
	GetTicketOrders(params domain.Params) ([]domain.TicketOrder, error)
	GetAttendeeByUserID(attendeeID int64) (queries.GetAttendeeByUserIDRow, error)
	GetTicketsByOrderID(orderID int64) ([]queries.GetTicketsByOrderIDRow, error)
}

type TicketPDFPort interface {
	GenerateTicket(queries.GetAttendeeByUserIDRow, []queries.GetTicketsByOrderIDRow) error
}

type TicketOrderRepo struct {
	db  TicketOrderRepoPort
	pdf TicketPDFPort
}

func NewTicketOrderRepo(db TicketOrderRepoPort, pdf TicketPDFPort) *TicketOrderRepo {
	return &TicketOrderRepo{db: db, pdf: pdf}
}

func (r *TicketOrderRepo) CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error) {
	ticketOrder, err := r.db.CreateTicketOrder(order, userID)
	if err != nil {
		return domain.TicketOrder{}, err
	}

	//attendee
	attendee, err := r.db.GetAttendeeByUserID(ticketOrder.AttendeeID)
	if err != nil {
		return domain.TicketOrder{}, err
	}

	//Get TicketOrderDetails
	tickets, err := r.db.GetTicketsByOrderID(ticketOrder.OrderID)
	if err != nil {
		return domain.TicketOrder{}, err
	}

	//Send To Ticket Generator
	r.pdf.GenerateTicket(attendee, tickets)

	return ticketOrder, nil
}

func (r *TicketOrderRepo) GetTicketOrders(params domain.Params) ([]domain.TicketOrder, error) {
	//Ticket Order
	torders, err := r.db.GetTicketOrders(params)

	return torders, err
}
