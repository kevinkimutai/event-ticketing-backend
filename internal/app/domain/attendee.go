package domain

type Attendee struct {
	AttendeeID int64 `json:"attendee_id"`
	UserID     int64 `json:"user_id"`
}

type AttendeeResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Attendee `json:"data"`
}
