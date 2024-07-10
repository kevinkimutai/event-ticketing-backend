package domain

type Category struct {
	CategoryID int64  `json:"category_id"`
	Name       string `json:"name"`
}

type CategoryResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Category `json:"data"`
}

type CategoriesResponse struct {
	StatusCode uint       `json:"status_code"`
	Message    string     `json:"message"`
	Data       []Category `json:"data"`
}
