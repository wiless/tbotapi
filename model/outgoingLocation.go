package model

import (
	"fmt"
)

type OutgoingLocation struct {
	OutgoingBase
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func NewOutgoingLocation(recipient Recipient, latitude, longitude float32) *OutgoingLocation {
	return &OutgoingLocation{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (ol *OutgoingLocation) GetQueryString() Querystring {
	toReturn := map[string]string(ol.GetBaseQueryString())

	toReturn["latitude"] = fmt.Sprint(ol.Latitude)
	toReturn["longitude"] = fmt.Sprint(ol.Longitude)

	return Querystring(toReturn)
}
