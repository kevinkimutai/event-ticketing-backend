package domain

type User struct {
	UserID   int64  `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       User   `json:"data"`
}
