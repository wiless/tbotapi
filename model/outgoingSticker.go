package model

import (
	"net/url"
)

type OutgoingSticker struct {
	OutgoingBase
}

func NewOutgoingSticker(chatId int) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (os *OutgoingSticker) GetQueryString() Querystring {
	toReturn := url.Values(os.GetBaseQueryString())

	return Querystring(toReturn)
}
