// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ticket_types.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTicketType = `-- name: CreateTicketType :one
INSERT INTO ticket_types (
  name,price,total_tickets,remaining_tickets,event_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING ticket_type_id, name, price, total_tickets, remaining_tickets, event_id
`

type CreateTicketTypeParams struct {
	Name             pgtype.Text
	Price            pgtype.Numeric
	TotalTickets     int32
	RemainingTickets int32
	EventID          int64
}

func (q *Queries) CreateTicketType(ctx context.Context, arg CreateTicketTypeParams) (TicketType, error) {
	row := q.db.QueryRow(ctx, createTicketType,
		arg.Name,
		arg.Price,
		arg.TotalTickets,
		arg.RemainingTickets,
		arg.EventID,
	)
	var i TicketType
	err := row.Scan(
		&i.TicketTypeID,
		&i.Name,
		&i.Price,
		&i.TotalTickets,
		&i.RemainingTickets,
		&i.EventID,
	)
	return i, err
}

const getEventTicketTypes = `-- name: GetEventTicketTypes :many
SELECT ticket_type_id, name, price, total_tickets, remaining_tickets, event_id FROM ticket_types
WHERE event_id = $1
`

func (q *Queries) GetEventTicketTypes(ctx context.Context, eventID int64) ([]TicketType, error) {
	rows, err := q.db.Query(ctx, getEventTicketTypes, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TicketType
	for rows.Next() {
		var i TicketType
		if err := rows.Scan(
			&i.TicketTypeID,
			&i.Name,
			&i.Price,
			&i.TotalTickets,
			&i.RemainingTickets,
			&i.EventID,
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

const getTicketType = `-- name: GetTicketType :one
SELECT ticket_type_id, name, price, total_tickets, remaining_tickets, event_id FROM ticket_types
WHERE ticket_type_id = $1 LIMIT 1
`

func (q *Queries) GetTicketType(ctx context.Context, ticketTypeID int64) (TicketType, error) {
	row := q.db.QueryRow(ctx, getTicketType, ticketTypeID)
	var i TicketType
	err := row.Scan(
		&i.TicketTypeID,
		&i.Name,
		&i.Price,
		&i.TotalTickets,
		&i.RemainingTickets,
		&i.EventID,
	)
	return i, err
}

const getTicketTypesByEvent = `-- name: GetTicketTypesByEvent :many
SELECT ticket_type_id, name, price, total_tickets, remaining_tickets, event_id FROM ticket_types
WHERE event_id = $1
`

func (q *Queries) GetTicketTypesByEvent(ctx context.Context, eventID int64) ([]TicketType, error) {
	rows, err := q.db.Query(ctx, getTicketTypesByEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TicketType
	for rows.Next() {
		var i TicketType
		if err := rows.Scan(
			&i.TicketTypeID,
			&i.Name,
			&i.Price,
			&i.TotalTickets,
			&i.RemainingTickets,
			&i.EventID,
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

const updateRemainingTicketType = `-- name: UpdateRemainingTicketType :exec
UPDATE ticket_types
  set remaining_tickets = $2

WHERE ticket_type_id = $1
`

type UpdateRemainingTicketTypeParams struct {
	TicketTypeID     int64
	RemainingTickets int32
}

func (q *Queries) UpdateRemainingTicketType(ctx context.Context, arg UpdateRemainingTicketTypeParams) error {
	_, err := q.db.Exec(ctx, updateRemainingTicketType, arg.TicketTypeID, arg.RemainingTickets)
	return err
}
