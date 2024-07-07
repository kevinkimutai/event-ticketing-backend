package domain

import "errors"

type TicketType struct {
	TicketTypeID int64
	Name         string
	Price        float64
	TotalTickets int32
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

	return nil
}
