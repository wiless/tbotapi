package model

import (
	"fmt"
)

type OutgoingVideo struct {
	OutgoingBase
	Duration int    `json:"duration,omitempty"`
	Caption  string `json:"caption,omitempty"`
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
	return ov
}

func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.Duration = to
	return ov
}

func (ov *OutgoingVideo) GetQueryString() Querystring {
	toReturn := map[string]string(ov.GetBaseQueryString())

	if ov.Caption != "" {
		toReturn["caption"] = ov.Caption
	}

	if ov.Duration != 0 {
		toReturn["duration"] = fmt.Sprint(ov.Duration)
	}

	return Querystring(toReturn)
}
