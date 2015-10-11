package model

import (
	"fmt"
)

type OutgoingAudioPub struct {
	OutgoingBasePub
	Duration  int    `json:"duration,omitempty"`
	Title     string `json:"title,omitempty"`
	Performer string `json:"performer,omitempty"`
}

type OutgoingAudio struct {
	OutgoingBase
	duration     int
	performer    string
	title        string
	durationSet  bool
	performerSet bool
	titleSet     bool
}

func NewOutgoingAudio(recipient Recipient) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (oa *OutgoingAudio) SetDuration(to int) *OutgoingAudio {
	oa.duration = to
	oa.durationSet = true
	return oa
}

func (oa *OutgoingAudio) SetPerformer(to string) *OutgoingAudio {
	oa.performer = to
	oa.performerSet = true
	return oa
}

func (oa *OutgoingAudio) SetTitle(to string) *OutgoingAudio {
	oa.title = to
	oa.titleSet = true
	return oa
}

func (oa *OutgoingAudio) GetQueryString() Querystring {
	toReturn := map[string]string(oa.GetBaseQueryString())

	if oa.durationSet {
		toReturn["duration"] = fmt.Sprint(oa.duration)
	}

	if oa.performerSet {
		toReturn["performer"] = oa.performer
	}

	if oa.titleSet {
		toReturn["title"] = oa.title
	}

	return Querystring(toReturn)
}

func (oa *OutgoingAudio) GetPub() OutgoingAudioPub {
	return OutgoingAudioPub{
		OutgoingBasePub: oa.GetPubBase(),
		Duration:        oa.duration,
		Performer:       oa.performer,
		Title:           oa.title,
	}
}
