package domain

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Attendee struct {
	AttendeeID int64 `json:"attendee_id"`
	UserID     int64 `json:"user_id"`
}

type AttendeeEvents struct {
	AttendeeID  int64       `json:"attendee_id"`
	EventName   string      `json:"event_name"`
	EventDate   time.Time   `json:"event_date"`
	PaymentID   pgtype.Int8 `json:"payment_id"`
	Quantity    int64       `json:"quantity"`
	TotalAmount float64     `json:"total_amount"`
}

type AttendeeData struct {
	TotalSpent  int64            `json:"total_spent"`
	TotalEvents int64            `json:"total_events"`
	Data        []AttendeeEvents `json:"data"`
}

type AttendeeResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Attendee `json:"data"`
}

type AttendeeEventFetch struct {
	Page          uint         `json:"page"`
	NumberOfPages uint         `json:"number_of_pages"`
	Total         int64        `json:"total"`
	Data          AttendeeData `json:"data"`
}

type AttendeesEventResponse struct {
	StatusCode    uint         `json:"status_code"`
	Message       string       `json:"message"`
	Page          uint         `json:"page"`
	NumberOfPages uint         `json:"number_of_pages"`
	Total         int64        `json:"total"`
	Data          AttendeeData `json:"data"`
}
