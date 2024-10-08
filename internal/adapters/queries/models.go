// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Admin struct {
	AdminID   int64
	UserID    int64
	CreatedAt pgtype.Timestamptz
}

type Attendee struct {
	AttendeeID int64
	UserID     int64
	OrderID    pgtype.Int8
	CreatedAt  pgtype.Timestamptz
}

type Category struct {
	CategoryID int64
	Name       string
}

type Event struct {
	EventID     int64
	Name        string
	CategoryID  int64
	Date        pgtype.Timestamptz
	FromTime    pgtype.Timestamptz
	ToTime      pgtype.Timestamptz
	Location    string
	Description pgtype.Text
	CreatedAt   pgtype.Timestamptz
	Longitude   float64
	Latitude    float64
	PosterUrl   string
	LocationID  int64
}

type Location struct {
	LocationID int64
	Name       string
}

type Organiser struct {
	OrganiserID int64
	UserID      int64
	EventID     int64
	CreatedAt   pgtype.Timestamptz
}

type Payment struct {
	PaymentID  int64
	StripeID   string
	Status     pgtype.Text
	TotalPrice pgtype.Numeric
}

type Ticket struct {
	TicketID     int64
	TicketTypeID int64
}

type TicketOrder struct {
	OrderID     int64
	PaymentID   pgtype.Int8
	CreatedAt   pgtype.Timestamptz
	AttendeeID  pgtype.Int8
	TotalAmount pgtype.Numeric
	AdmitStatus pgtype.Bool
}

type TicketOrderItem struct {
	ItemID     int64
	OrderID    int64
	TicketID   int64
	Quantity   int64
	TotalPrice pgtype.Numeric
}

type TicketType struct {
	TicketTypeID     int64
	Name             pgtype.Text
	Price            pgtype.Numeric
	TotalTickets     int32
	RemainingTickets int32
	EventID          int64
}

type User struct {
	UserID    int64
	FullName  string
	Email     string
	CreatedAt pgtype.Timestamptz
}

type Usher struct {
	UsherID   int64
	UserID    int64
	CreatedAt pgtype.Timestamptz
}
