package db

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/internal/adapters/queries"
)

func (db *DBAdapter) GetTicketsByOrderID(orderID int64) ([]queries.GetTicketsByOrderIDRow, error) {
	tickets, err := db.queries.GetTicketsByOrderID(context.Background(), orderID)

	return tickets, err
}
