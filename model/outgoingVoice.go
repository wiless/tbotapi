package model

import (
	"fmt"
	"net/url"
)

type OutgoingVoicePub struct {
	OutgoingBasePub
	Duration int `json:"duration,omitempty"`
}

type OutgoingVoice struct {
	OutgoingBase
	duration    int
	durationSet bool
}

func NewOutgoingVoice(chatId int) *OutgoingVoice {
	return &OutgoingVoice{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
	}
}

func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := url.Values(ov.GetBaseQueryString())

	if ov.durationSet {
		toReturn.Set("duration", fmt.Sprint(ov.duration))
	}

	return Querystring(toReturn)
}

func (ov *OutgoingVoice) GetPub() OutgoingVoicePub {
	return OutgoingVoicePub{
		OutgoingBasePub: ov.GetPubBase(),
		Duration:        ov.duration,
	}
}
