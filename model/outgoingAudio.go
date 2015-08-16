package model

import (
	"fmt"
	"net/url"
)

type OutgoingAudio struct {
	OutgoingBase
	duration     int
	performer    string
	title        string
	durationSet  bool
	performerSet bool
	titleSet     bool
}

func NewOutgoingAudio(chatId int) *OutgoingAudio {
	return &OutgoingAudio{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
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
	toReturn := url.Values(oa.GetBaseQueryString())

	if oa.durationSet {
		toReturn.Set("duration", fmt.Sprint(oa.duration))
	}

	if oa.performerSet {
		toReturn.Set("performer", oa.performer)
	}

	if oa.titleSet {
		toReturn.Set("title", oa.title)
	}

	return Querystring(toReturn)
}
