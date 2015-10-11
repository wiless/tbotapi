package model

import (
	"fmt"
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

func NewOutgoingVoice(recipient Recipient) *OutgoingVoice {
	return &OutgoingVoice{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.durationSet {
		toReturn["duration"] = fmt.Sprint(ov.duration)
	}

	return Querystring(toReturn)
}

func (ov *OutgoingVoice) GetPub() OutgoingVoicePub {
	return OutgoingVoicePub{
		OutgoingBasePub: ov.GetPubBase(),
		Duration:        ov.duration,
	}
}
