package model

import (
	"fmt"
)

// OutgoingAudio represents an outgoing audio file
type OutgoingAudio struct {
	OutgoingBase
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

// NewOutgoingAudio creates a new outgoing audio file
func NewOutgoingAudio(recipient Recipient) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetDuration sets a duration for the audio file (optional)
func (oa *OutgoingAudio) SetDuration(to int) *OutgoingAudio {
	oa.Duration = to
	return oa
}

// SetPerformer sets a performer for the audio file (optional)
func (oa *OutgoingAudio) SetPerformer(to string) *OutgoingAudio {
	oa.Performer = to
	return oa
}

// SetTitle sets a title for the audio file (optional)
func (oa *OutgoingAudio) SetTitle(to string) *OutgoingAudio {
	oa.Title = to
	return oa
}

// GetQueryString returns a Querystring representing the audio file
func (oa *OutgoingAudio) GetQueryString() Querystring {
	toReturn := map[string]string(oa.GetBaseQueryString())

	if oa.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(oa.Duration)
	}

	if oa.Performer != "" {
		toReturn["performer"] = oa.Performer
	}

	if oa.Title != "" {
		toReturn["title"] = oa.Title
	}

	return Querystring(toReturn)
}
