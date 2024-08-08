package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
	"github.com/kevinkimutai/ticketingapp/internal/app/domain"
)

func (db *DBAdapter) GetTicketsByOrderID(orderID int64) ([]queries.GetTicketsByOrderIDRow, error) {
	tickets, err := db.queries.GetTicketsByOrderID(context.Background(), orderID)

	return tickets, err
}

func (db *DBAdapter) GetTicket(ticketTypeID int64) (domain.Ticket, error) {
	ctx := context.Background()

	ticket, err := db.queries.GetTicketByTicketTypeID(ctx, ticketTypeID)

	return domain.Ticket{
		TicketID:     ticket.TicketID,
		TicketTypeID: ticket.TicketTypeID,
	}, err
}
