package model

// UserResponse represents the response sent by the API on a GetMe request
type UserResponse struct {
	BaseResponse
	User User `json:"result"`
}

// User represents a Telegram user or bot
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
