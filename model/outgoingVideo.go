package model

import (
	"fmt"
	"net/url"
)

type OutgoingVideo struct {
	OutgoingBase
	duration    int
	durationSet bool
	caption     string
	captionSet  bool
}

func NewOutgoingVideo(chatId int) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (ov *OutgoingVideo) SetCaption(to string) *OutgoingVideo {
	ov.caption = to
	ov.captionSet = true
	return ov
}

func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := url.Values(ov.GetBaseQueryString())

	if ov.captionSet {
		toReturn.Set("caption", ov.caption)
	}

	if ov.durationSet {
		toReturn.Set("duration", fmt.Sprint(ov.duration))
	}

	return Querystring(toReturn)
}
