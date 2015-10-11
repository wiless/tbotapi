package model

import (
	"fmt"
)

type OutgoingVideoPub struct {
	OutgoingBasePub
	Duration int    `json:"duration,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

type OutgoingVideo struct {
	OutgoingBase
	duration    int
	durationSet bool
	caption     string
	captionSet  bool
}

func NewOutgoingVideo(recipient Recipient) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (ov *OutgoingVideo) SetCaption(to string) *OutgoingVideo {
	ov.caption = to
	ov.captionSet = true
	return ov
}

func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.captionSet {
		toReturn["caption"] = ov.caption
	}

	if ov.durationSet {
		toReturn["duration"] = fmt.Sprint(ov.duration)
	}

	return Querystring(toReturn)
}

func (ov *OutgoingVideo) GetPub() OutgoingVideoPub {
	return OutgoingVideoPub{
		OutgoingBasePub: ov.GetPubBase(),
		Duration:        ov.duration,
		Caption:         ov.caption,
	}
}
