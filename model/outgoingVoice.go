package model

import (
	"fmt"
)

type OutgoingVoice struct {
	OutgoingBase
	Duration    int `json:"duration,omitempty"`
	durationSet bool
}

func NewOutgoingVoice(recipient Recipient) *OutgoingVoice {
	return &OutgoingVoice{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.Duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.durationSet {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}
