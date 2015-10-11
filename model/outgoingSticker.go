package model

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
	return os.GetBaseQueryString()
}

func (os *OutgoingSticker) GetPub() OutgoingStickerPub {
	return OutgoingStickerPub{
		OutgoingBasePub: os.GetPubBase(),
	}
}
