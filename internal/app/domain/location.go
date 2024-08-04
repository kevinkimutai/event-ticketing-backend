package domain

type Location struct {
	LocationID int64  `json:"location_id"`
	Name       string `json:"name"`
}

type LocationResponse struct {
	StatusCode uint     `json:"status_code"`
	Message    string   `json:"message"`
	Data       Location `json:"data"`
}

// type CategoriesResponse struct {
// 	StatusCode uint       `json:"status_code"`
// 	Message    string     `json:"message"`
// 	Data       []Category `json:"data"`
// }
