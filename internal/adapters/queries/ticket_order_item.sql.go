// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ticket_order_item.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTicketOrderItem = `-- name: CreateTicketOrderItem :one
INSERT INTO ticket_order_items (
  order_id,ticket_id,quantity,total_price
) VALUES (
  $1, $2, $3, $4
)
RETURNING item_id, order_id, ticket_id, quantity, total_price
`

type CreateTicketOrderItemParams struct {
	OrderID    int64
	TicketID   int64
	Quantity   int64
	TotalPrice pgtype.Numeric
}

func (q *Queries) CreateTicketOrderItem(ctx context.Context, arg CreateTicketOrderItemParams) (TicketOrderItem, error) {
	row := q.db.QueryRow(ctx, createTicketOrderItem,
		arg.OrderID,
		arg.TicketID,
		arg.Quantity,
		arg.TotalPrice,
	)
	var i TicketOrderItem
	err := row.Scan(
		&i.ItemID,
		&i.OrderID,
		&i.TicketID,
		&i.Quantity,
		&i.TotalPrice,
	)
	return i, err
}

const getTicketOrderItemByTicketID = `-- name: GetTicketOrderItemByTicketID :one
SELECT item_id, order_id, ticket_id, quantity, total_price FROM ticket_order_items
WHERE ticket_id = $1 
LIMIT 1
`

func (q *Queries) GetTicketOrderItemByTicketID(ctx context.Context, ticketID int64) (TicketOrderItem, error) {
	row := q.db.QueryRow(ctx, getTicketOrderItemByTicketID, ticketID)
	var i TicketOrderItem
	err := row.Scan(
		&i.ItemID,
		&i.OrderID,
		&i.TicketID,
		&i.Quantity,
		&i.TotalPrice,
	)
	return i, err
}
