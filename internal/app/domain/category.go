package domain

type Category struct {
	CategoryID int64
	Name       string
}

type CategoryResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Category `json:"data"`
}
