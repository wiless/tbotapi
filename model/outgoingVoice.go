package model

import (
	"fmt"
)

type OutgoingVoice struct {
	OutgoingBase
	Duration int `json:"duration,omitempty"`
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
	return ov
}

func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}
