package model

import (
	"net/url"
)

type OutgoingAudio struct {
	OutgoingBase
}

func NewOutgoingAudio(chatId int) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (oa *OutgoingAudio) GetQueryString() Querystring {
	toReturn := url.Values(oa.GetBaseQueryString())

	return Querystring(toReturn)
}
