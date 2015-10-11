package model

type OutgoingSticker struct {
	OutgoingBase
}

func NewOutgoingSticker(recipient Recipient) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (os *OutgoingSticker) GetQueryString() Querystring {
	return os.GetBaseQueryString()
}
