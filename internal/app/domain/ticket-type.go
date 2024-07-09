package domain

import "errors"

type TicketType struct {
	EventID          int64   `json:"event_id"`
	TicketTypeID     int64   `json:"ticket_type_id"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	TotalTickets     int32   `json:"total_tickets"`
	RemainingTickets int32   `json:"remaining_tickets"`
}

type TicketTypeResponse struct {
	StatusCode uint       `json:"status_code"`
	Message    string     `json:"message"`
	Data       TicketType `json:"data"`
}

func NewTicketTypeDomain(t *TicketType) error {

	if t.Name == "" {
		return errors.New("missing name field")
	}
	if t.Price == 0 {
		return errors.New("missing price field")
	}
	if t.TotalTickets == 0 {
		return errors.New("missing totaltickets field")
	}
	if t.RemainingTickets == 0 {
		t.RemainingTickets = t.TotalTickets
	}

	return nil
}
