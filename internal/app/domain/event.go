package domain

import (
	"errors"
	"strconv"
	"time"
)

type Event struct {
	EventID     int64     `json:"event_id"`
	Name        string    `json:"name"`
	CategoryID  int64     `json:"category_id"`
	Date        time.Time `json:"date"`
	FromTime    time.Time `json:"from_time"`
	ToTime      time.Time `json:"to_time"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	CreatedAt   time.Time `json:"created_at"`
	PosterUrl   string    `json:"poster_url"`
}

type EventsFetch struct {
	Page          uint    `json:"page"`
	NumberOfPages uint    `json:"number_of_pages"`
	Total         uint    `json:"total"`
	Data          []Event `json:"data"`
}

type EventsResponse struct {
	StatusCode    uint    `json:"status_code"`
	Message       string  `json:"message"`
	Page          uint    `json:"page"`
	NumberOfPages uint    `json:"number_of_pages"`
	Total         uint    `json:"total"`
	Data          []Event `json:"data"`
}

type EventResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       Event  `json:"data"`
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

type Params struct {
	Page     int32
	Limit    int32
	Category string
}

//Check Missing Data

func NewEventDomain(e *Event) error {
	if e.Name == "" {
		return errors.New("missing name field")
	}
	if e.CategoryID == 0 {
		return errors.New("missing category field")
	}
	if e.Date.String() == "" {
		return errors.New("missing date field")
	}
	if e.FromTime.String() == "" {
		return errors.New("missing fromtime field")
	}
	if e.ToTime.String() == "" {
		return errors.New("missing totime field")
	}
	if e.Description == "" {
		return errors.New("missing totime field")
	}
	if e.Location == "" {
		return errors.New("missing totime field")
	}
	if e.Longitude == 0 {
		return errors.New("missing totime field")
	}
	if e.Latitude == 0 {
		return errors.New("missing totime field")
	}
	if e.PosterUrl == "" {
		return errors.New("missing poster")
	}

	return nil
}

func CheckEventParams(m map[string]string) Params {
	var LIMIT, OFFSET int32 = 10, 0

	if m["limit"] != "" {
		items, _ := strconv.Atoi(m["limit"])

		LIMIT = int32(items)

	}
	if m["page"] != "" {
		page, _ := strconv.Atoi(m["page"])

		if page < 1 {
			page = 1
		}

		OFFSET = (int32(page) - 1) * LIMIT

	}

	return Params{
		Page:     OFFSET,
		Limit:    LIMIT,
		Category: m["category"],
	}
}
