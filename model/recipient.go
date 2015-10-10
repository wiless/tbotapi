package model

import "fmt"

type Recipient struct {
	ChatId    *int
	ChannelId *string
}

func NewChatRecipient(chatId int) Recipient {
	return Recipient{
		ChatId: &chatId,
	}
}

func NewChannelRecipient(channelName string) Recipient {
	return Recipient{
		ChannelId: &channelName,
	}
}

func NewRecipientFromChat(chat Chat) Recipient {
	return NewChatRecipient(chat.Id) //No need to distinguish between channels and chats, bots cannot receive from channels
}

func (r Recipient) isChat() bool {
	return r.ChatId != nil
}

func (r Recipient) isChannel() bool {
	return r.ChannelId != nil
}

func (r Recipient) MarshalJSON() ([]byte, error) {
	toReturn := ""

	if r.isChannel() {
		toReturn = fmt.Sprintf("\"%s\"", *r.ChannelId)
	} else {
		toReturn = fmt.Sprintf("%d", *r.ChatId)
	}

	return []byte(toReturn), nil
}
