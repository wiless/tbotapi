package model

import (
	"fmt"
)

type OutgoingAudio struct {
	OutgoingBase
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

func NewOutgoingAudio(recipient Recipient) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (oa *OutgoingAudio) SetDuration(to int) *OutgoingAudio {
	oa.Duration = to
	return oa
}

func (oa *OutgoingAudio) SetPerformer(to string) *OutgoingAudio {
	oa.Performer = to
	return oa
}

func (oa *OutgoingAudio) SetTitle(to string) *OutgoingAudio {
	oa.Title = to
	return oa
}

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
