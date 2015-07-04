package model

import (
	"net/url"
)

type OutgoingVideo struct {
	OutgoingBase
}

func NewOutgoingVideo(chatId int) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := url.Values(ov.GetBaseQueryString())

	return Querystring(toReturn)
}
