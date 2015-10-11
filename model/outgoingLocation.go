package model

import (
	"fmt"
)

// OutgoingLocation represents an outgoing location on a map
type OutgoingLocation struct {
	OutgoingBase
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// NewOutgoingLocation creates a new outgoing location
func NewOutgoingLocation(recipient Recipient, latitude, longitude float32) *OutgoingLocation {
	return &OutgoingLocation{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// GetQueryString returns a Querystring representing the location
func (ol *OutgoingLocation) GetQueryString() Querystring {
	toReturn := map[string]string(ol.GetBaseQueryString())

	toReturn["latitude"] = fmt.Sprint(ol.Latitude)
	toReturn["longitude"] = fmt.Sprint(ol.Longitude)

	return Querystring(toReturn)
}
