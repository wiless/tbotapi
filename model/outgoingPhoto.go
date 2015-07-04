package model

import (
	"net/url"
)

type OutgoingPhoto struct {
	OutgoingBase
	caption    string
	captionSet bool
}

func NewOutgoingPhoto(chatId int) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.caption = to
	op.captionSet = true
	return op
}

func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := url.Values(op.GetBaseQueryString())

	if op.captionSet {
		toReturn.Set("caption", op.caption)
	}

	return Querystring(toReturn)
}
