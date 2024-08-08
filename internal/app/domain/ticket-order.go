package domain

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Ticket struct {
	TicketID     int64
	EventID      int64
	TicketTypeID int64
}

type TicketOrder struct {
	OrderID     int64       `json:"order_id"`
	PaymentID   pgtype.Int8 `json:"payment_id"`
	CreatedAt   time.Time   `json:"created_at"`
	AttendeeID  int64       `json:"attendee_id"`
	TotalAmount float64     `json:"total_amount"`
}

type TicketOrderItem struct {
	ItemID       int64   `json:"item_id"`
	OrderID      int64   `json:"order_id"`
	TicketTypeID int64   `json:"ticket_type_id"`
	Quantity     int64   `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
}

type OrderItem struct {
	ItemID     int64
	OrderID    int64
	TicketID   int64
	Quantity   int64
	TotalPrice float64
}

type TicketOrderResponse struct {
	StatusCode uint        `json:"status_code"`
	Message    string      `json:"message"`
	Data       TicketOrder `json:"data"`
}
type TicketOrdersResponse struct {
	StatusCode uint          `json:"status_code"`
	Message    string        `json:"message"`
	Data       []TicketOrder `json:"data"`
}

type TicketResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       Ticket `json:"data"`
}
type TicketOrderRequest struct {
	OrderItems []TicketOrderItem `json:"order_items"`
}
