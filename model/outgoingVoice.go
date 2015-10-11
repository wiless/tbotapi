package model

import (
	"fmt"
)

// OutgoingVoice represents an outgoing voice note
type OutgoingVoice struct {
	OutgoingBase
	Duration int `json:"duration,omitempty"`
}

// NewOutgoingVoice creates a new outgoing voice note
func NewOutgoingVoice(recipient Recipient) *OutgoingVoice {
	return &OutgoingVoice{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetDuration sets a duration of the voice note (optional)
func (ov *OutgoingVoice) SetDuration(to int) *OutgoingVoice {
	ov.Duration = to
	return ov
}

// GetQueryString returns a Querystring representing the outgoing voice note
func (ov *OutgoingVoice) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}
