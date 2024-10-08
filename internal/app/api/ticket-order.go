package api

import (
	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

type TicketOrderRepoPort interface {
	CreateTicketOrder(order *domain.TicketOrderRequest, userID int64) (domain.TicketOrder, error)
	GetTicketOrders(params *domain.Params) ([]domain.TicketOrder, error)
	GetAttendeeByUserID(attendeeID int64) (queries.GetAttendeeByUserIDRow, error)
	GetTicketsByOrderID(orderID int64) ([]queries.GetTicketsByOrderIDRow, error)
	GetOrderItemByTicketID(ticketID int64) (domain.TicketOrderItem, error)

	GetTicketOrder(orderItemId int64) (domain.TicketOrder, error)
	GetTicketOrderDetails(orderID int64) (domain.TicketOrderDetails, error)
	AdmitTicketOrder(orderId int64) error
}

type TicketPDFPort interface {
	GenerateTicket(queries.GetAttendeeByUserIDRow, []queries.GetTicketsByOrderIDRow) (string, error)
}

type QueuePort interface {
	SendOrderConfirmation(email, fullName, ticketPDFUrl string)
}

type TicketOrderRepo struct {
	db    TicketOrderRepoPort
	pdf   TicketPDFPort
	queue QueuePort
}

func NewTicketOrderRepo(db TicketOrderRepoPort, pdf TicketPDFPort, queue QueuePort) *TicketOrderRepo {
	return &TicketOrderRepo{
		db:    db,
		pdf:   pdf,
		queue: queue}
}

func (r *TicketOrderRepo) CreateTicketOrder(order *domain.TicketOrderRequest, user queries.User) (domain.TicketOrder, error) {
	ticketOrder, err := r.db.CreateTicketOrder(order, user.UserID)
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
	filePath, err := r.pdf.GenerateTicket(attendee, tickets)
	if err != nil {
		return domain.TicketOrder{}, err
	}

	//Send Email Queue
	r.queue.SendOrderConfirmation(user.Email, user.FullName, filePath)

	return ticketOrder, nil
}

func (r *TicketOrderRepo) GetTicketOrders(params *domain.Params) ([]domain.TicketOrder, error) {
	//Ticket Order
	torders, err := r.db.GetTicketOrders(params)

	return torders, err
}

func (r *TicketOrderRepo) GetOrderItemByTicketID(ticketID int64) (domain.TicketOrderItem, error) {
	//Ticket OrderItem
	torderitem, err := r.db.GetOrderItemByTicketID(ticketID)

	return torderitem, err
}

func (r *TicketOrderRepo) GetOrder(orderItemID int64) (domain.TicketOrder, error) {
	tOrder, err := r.db.GetTicketOrder(orderItemID)

	return tOrder, err
}

func (r *TicketOrderRepo) GetTicketOrderDetails(orderID int64) (domain.TicketOrderDetails, error) {
	tOrder, err := r.db.GetTicketOrderDetails(orderID)

	return tOrder, err
}
func (r *TicketOrderRepo) AdmitTicketOrder(orderId int64) error {
	err := r.db.AdmitTicketOrder(orderId)

	return err
}
