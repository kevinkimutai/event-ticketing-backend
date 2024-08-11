package domain

import "time"

type Organiser struct {
	OrganiserID int64     `json:"organiser_id"`
	UserID      int64     `json:"user_id"`
	EventID     int64     `json:"event_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrganisersFetch struct {
	Page          uint        `json:"page"`
	NumberOfPages uint        `json:"number_of_pages"`
	Total         int64       `json:"total"`
	Data          []Organiser `json:"data"`
}

type OrganisersResponse struct {
	StatusCode    uint        `json:"status_code"`
	Message       string      `json:"message"`
	Page          uint        `json:"page"`
	NumberOfPages uint        `json:"number_of_pages"`
	Total         int64       `json:"total"`
	Data          []Organiser `json:"data"`
}
