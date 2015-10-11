package model

import (
	"fmt"
)

type OutgoingVideo struct {
	OutgoingBase
	Duration    int    `json:"duration,omitempty"`
	Caption     string `json:"caption,omitempty"`
	durationSet bool
	captionSet  bool
}

func NewOutgoingVideo(recipient Recipient) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (ov *OutgoingVideo) SetCaption(to string) *OutgoingVideo {
	ov.Caption = to
	ov.captionSet = true
	return ov
}

func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.Duration = to
	ov.durationSet = true
	return ov
}

func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.captionSet {
		toReturn["caption"] = ov.Caption
	}

	if ov.durationSet {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}
