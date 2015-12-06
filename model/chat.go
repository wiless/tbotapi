package model

import (
	"fmt"
)

// Chat contains information about the chat a message originated from
type Chat struct {
	ID        int     `json:"id"`         // Unique identifier for this chat
	Type      string  `json:"type"`       // Type of chat, can be either "private", "group" or "channel". Check Is(PrivateChat|GroupChat|Channel)() methods
	Title     *string `json:"title"`      // Title for channels and group chats
	Username  *string `json:"username"`   // Username for private chats and channels if available
	FirstName *string `json:"first_name"` // First name of the other party in a private chat
	LastName  *string `json:"last_name"`  // Last name of the other party in a private chat
}

// IsPrivateChat checks if the chat is a private chat
func (c Chat) IsPrivateChat() bool {
	return c.Type == "private"
}

// IsGroupChat checks if the chat is a group chat
func (c Chat) IsGroupChat() bool {
	return c.Type == "group"
}

// IsChannel checks if the chat is a channel
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

func (c Chat) String() string {
	toReturn := fmt.Sprint(c.ID)
	if c.IsPrivateChat() {
		toReturn += " (P) "
	} else if c.IsGroupChat() {
		toReturn += " (G) "
	} else {
		toReturn += " (C) "
	}

	if c.Title != nil {
		toReturn += "\"" + *c.Title + "\" "
	}

	if c.FirstName != nil {
		toReturn += *c.FirstName + " "
	}

	if c.LastName != nil {
		toReturn += *c.LastName + " "
	}

	if c.Username != nil {
		toReturn += "(@" + *c.Username + ")"
	}

	return toReturn
}
