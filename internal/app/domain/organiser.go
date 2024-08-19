package domain

import (
	"time"
)

type OrganiserData struct {
	TotalAmountEvents float64     `json:"total_amount_events"`
	Data              []Organiser `json:"data"`
}
type Organiser struct {
	OrganiserID int64     `json:"organiser_id"`
	UserID      int64     `json:"user_id"`
	EventID     int64     `json:"event_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrganisersFetch struct {
	Page          uint          `json:"page"`
	NumberOfPages uint          `json:"number_of_pages"`
	Total         int64         `json:"total"`
	Data          OrganiserData `json:"data"`
}

type OrganisersResponse struct {
	StatusCode    uint          `json:"status_code"`
	Message       string        `json:"message"`
	Page          uint          `json:"page"`
	NumberOfPages uint          `json:"number_of_pages"`
	Total         int64         `json:"total"`
	Data          OrganiserData `json:"data"`
}

type OrganiserEvent struct {
	AttendeeID     int64   `json:"attendee_id"`
	Fullname       string  `json:"full_name"`
	Email          string  `json:"email"`
	TicketTypeName string  `json:"ticket_type_name"`
	Quantity       int64   `json:"quantity"`
	Total          float64 `json:"total"`
}

type OrganiserEventFetch struct {
	Page          uint             `json:"page"`
	NumberOfPages uint             `json:"number_of_pages"`
	Total         int64            `json:"total"`
	TicketsSold   float64          `json:"tickets_sold"`
	TotalAmount   float64          `json:"total_amount"`
	Data          []OrganiserEvent `json:"data"`
}

type OrganiserEventResponse struct {
	StatusCode    uint             `json:"status_code"`
	Message       string           `json:"message"`
	Page          uint             `json:"page"`
	NumberOfPages uint             `json:"number_of_pages"`
	Total         int64            `json:"total"`
	TicketsSold   float64          `json:"tickets_sold"`
	TotalAmount   float64          `json:"total_amount"`
	Data          []OrganiserEvent `json:"data"`
}
