// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ticket.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets (
  ticket_type_id
) VALUES (
  $1
)
RETURNING ticket_id, ticket_type_id
`

func (q *Queries) CreateTicket(ctx context.Context, ticketTypeID int64) (Ticket, error) {
	row := q.db.QueryRow(ctx, createTicket, ticketTypeID)
	var i Ticket
	err := row.Scan(&i.TicketID, &i.TicketTypeID)
	return i, err
}

const getTicketsByOrderID = `-- name: GetTicketsByOrderID :many
SELECT item_id, order_id, it.ticket_id, quantity, total_price, t.ticket_id, t.ticket_type_id, tty.ticket_type_id, tty.name, price, total_tickets, remaining_tickets, tty.event_id, ev.event_id, ev.name, category_id, date, from_time, to_time, location, description, created_at, longitude, latitude, poster_url FROM ticket_order_items it
JOIN tickets t
ON t.ticket_id = it.ticket_id
JOIN ticket_types tty
ON tty.ticket_type_id=t.ticket_type_id
JOIN events ev
ON ev.event_id =tty.event_id
WHERE it.order_id = $1
`

type GetTicketsByOrderIDRow struct {
	ItemID           int64
	OrderID          int64
	TicketID         int64
	Quantity         int64
	TotalPrice       pgtype.Numeric
	TicketID_2       int64
	TicketTypeID     int64
	TicketTypeID_2   int64
	Name             pgtype.Text
	Price            pgtype.Numeric
	TotalTickets     int32
	RemainingTickets int32
	EventID          int64
	EventID_2        int64
	Name_2           string
	CategoryID       int64
	Date             pgtype.Timestamptz
	FromTime         pgtype.Timestamptz
	ToTime           pgtype.Timestamptz
	Location         string
	Description      pgtype.Text
	CreatedAt        pgtype.Timestamptz
	Longitude        float64
	Latitude         float64
	PosterUrl        string
}

func (q *Queries) GetTicketsByOrderID(ctx context.Context, orderID int64) ([]GetTicketsByOrderIDRow, error) {
	rows, err := q.db.Query(ctx, getTicketsByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTicketsByOrderIDRow
	for rows.Next() {
		var i GetTicketsByOrderIDRow
		if err := rows.Scan(
			&i.ItemID,
			&i.OrderID,
			&i.TicketID,
			&i.Quantity,
			&i.TotalPrice,
			&i.TicketID_2,
			&i.TicketTypeID,
			&i.TicketTypeID_2,
			&i.Name,
			&i.Price,
			&i.TotalTickets,
			&i.RemainingTickets,
			&i.EventID,
			&i.EventID_2,
			&i.Name_2,
			&i.CategoryID,
			&i.Date,
			&i.FromTime,
			&i.ToTime,
			&i.Location,
			&i.Description,
			&i.CreatedAt,
			&i.Longitude,
			&i.Latitude,
			&i.PosterUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
