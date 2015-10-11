package model

type OutgoingForward struct {
	OutgoingBase
	FromChatID Recipient `json:"from_chat_id"`
	MessageID  int       `json:"message_id"`
}

func NewOutgoingForward(recipient Recipient, origin Chat, messageId int) *OutgoingForward {
	return &OutgoingForward{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		FromChatID: NewRecipientFromChat(origin),
		MessageID:  messageId,
	}
}
