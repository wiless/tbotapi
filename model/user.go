package model

import "fmt"

// UserResponse represents the response sent by the API on a GetMe request
type UserResponse struct {
	BaseResponse
	User User `json:"result"`
}

// User represents a Telegram user or bot
type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
}

func (u User) String() string {
	if u.LastName != nil && u.Username != nil {
		return fmt.Sprintf("%d/%s %s (@%s)", u.ID, u.FirstName, u.LastName, u.Username)
	} else if u.LastName != nil {
		return fmt.Sprintf("%d/%s %s", u.ID, u.FirstName, u.LastName)
	} else if u.Username != nil {
		return fmt.Sprintf("%d/%s (@%s)", u.ID, u.FirstName, u.Username)
	}
	return fmt.Sprintf("%d/%s", u.ID, u.FirstName)
}
