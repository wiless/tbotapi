package model

import (
	"net/url"
)

type OutgoingStickerPub struct {
	OutgoingBasePub
}

type OutgoingSticker struct {
	OutgoingBase
}

func NewOutgoingSticker(recipient Recipient) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (os *OutgoingSticker) GetQueryString() Querystring {
	toReturn := url.Values(os.GetBaseQueryString())

	return Querystring(toReturn)
}

func (os *OutgoingSticker) GetPub() OutgoingStickerPub {
	return OutgoingStickerPub{
		OutgoingBasePub: os.GetPubBase(),
	}
}
