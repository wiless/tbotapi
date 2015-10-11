package model

import (
	"fmt"
)

type OutgoingAudio struct {
	OutgoingBase
	Duration     int    `json:"duration,omitempty"`
	Title        string `json:"title,omitempty"`
	Performer    string `json:"performer,omitempty"`
	durationSet  bool
	performerSet bool
	titleSet     bool
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
	oa.durationSet = true
	return oa
}

func (oa *OutgoingAudio) SetPerformer(to string) *OutgoingAudio {
	oa.Performer = to
	oa.performerSet = true
	return oa
}

func (oa *OutgoingAudio) SetTitle(to string) *OutgoingAudio {
	oa.Title = to
	oa.titleSet = true
	return oa
}

func (oa *OutgoingAudio) GetQueryString() Querystring {
	toReturn := map[string]string(oa.GetBaseQueryString())

	if oa.durationSet {
		toReturn["duration"] = fmt.Sprint(oa.Duration)
	}

	if oa.performerSet {
		toReturn["performer"] = oa.Performer
	}

	if oa.titleSet {
		toReturn["title"] = oa.Title
	}

	return Querystring(toReturn)
}
