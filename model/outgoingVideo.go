package model

import (
	"fmt"
)

// OutgoingVideo represents an outgoing video file
type OutgoingVideo struct {
	OutgoingBase
	Duration int    `json:"duration,omitempty"`
	Caption  string `json:"caption,omitempty"`
}

// NewOutgoingVideo creates a new outgoing video file
func NewOutgoingVideo(recipient Recipient) *OutgoingVideo {
	return &OutgoingVideo{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetCaption sets a caption for the video file (optional)
func (ov *OutgoingVideo) SetCaption(to string) *OutgoingVideo {
	ov.Caption = to
	return ov
}

// SetDuration sets a duration for the video file (optional)
func (ov *OutgoingVideo) SetDuration(to int) *OutgoingVideo {
	ov.Duration = to
	return ov
}

// GetQueryString returns a Querystring representing the outgoing video file
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
