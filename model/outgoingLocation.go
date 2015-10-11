package model

import (
	"fmt"
)

type OutgoingLocationPub struct {
	OutgoingBasePub
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type OutgoingLocation struct {
	OutgoingBase
	latitude  float32
	longitude float32
}

func NewOutgoingLocation(recipient Recipient, latitude, longitude float32) *OutgoingLocation {
	return &OutgoingLocation{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
		latitude:  latitude,
		longitude: longitude,
	}
}

func (ol *OutgoingLocation) GetQueryString() Querystring {
	toReturn := map[string]string(ol.GetBaseQueryString())

	toReturn["latitude"] = fmt.Sprint(ol.latitude)
	toReturn["longitude"] = fmt.Sprint(ol.longitude)

	return Querystring(toReturn)
}

func (ol *OutgoingLocation) GetPub() OutgoingLocationPub {
	return OutgoingLocationPub{
		OutgoingBasePub: ol.GetPubBase(),
		Latitude:        ol.latitude,
		Longitude:       ol.longitude,
	}
}
