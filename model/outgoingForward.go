package model

// OutgoingForward represents an outgoing, forwarded message
type OutgoingForward struct {
	OutgoingBase
	FromChatID Recipient `json:"from_chat_id"`
	MessageID  int       `json:"message_id"`
}

// NewOutgoingForward creates a new outgoing, forwarded message
func NewOutgoingForward(recipient Recipient, origin Chat, messageID int) *OutgoingForward {
	return &OutgoingForward{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		FromChatID: NewRecipientFromChat(origin),
		MessageID:  messageID,
	}
}
