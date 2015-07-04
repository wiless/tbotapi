package model

import (
	"fmt"
	"net/url"
)

type OutgoingLocation struct {
	OutgoingBase
	latitude  float32
	longitude float32
}

func NewOutgoingLocation(chatId int, latitude, longitude float32) *OutgoingLocation {
	return &OutgoingLocation{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
		latitude:  latitude,
		longitude: longitude,
	}
}

func (ol *OutgoingLocation) GetQueryString() Querystring {
	toReturn := url.Values(ol.GetBaseQueryString())

	toReturn.Set("latitude", fmt.Sprint(ol.latitude))
	toReturn.Set("longitude", fmt.Sprint(ol.longitude))

	return Querystring(toReturn)
}
