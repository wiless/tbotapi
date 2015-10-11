package model

type OutgoingForwardPub struct {
	OutgoingBasePub
	FromChatId Recipient `json:"from_chat_id"`
	MessageId  int       `json:"message_id"`
}

type OutgoingForward struct {
	OutgoingBase
	from      Recipient
	messageId int
}

func NewOutgoingForward(recipient Recipient, origin Chat, messageId int) *OutgoingForward {
	return &OutgoingForward{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
		from:      NewRecipientFromChat(origin),
		messageId: messageId,
	}
}

func (of *OutgoingForward) GetPub() OutgoingForwardPub {
	return OutgoingForwardPub{
		OutgoingBasePub: of.GetPubBase(),
		FromChatId:      of.from,
		MessageId:       of.messageId,
	}
}
