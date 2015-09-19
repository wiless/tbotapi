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

func NewOutgoingSticker(chatId int) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
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
