package model

type OutgoingPhotoPub struct {
	OutgoingBasePub
	Caption string `json:"caption"`
}

type OutgoingPhoto struct {
	OutgoingBase
	caption    string
	captionSet bool
}

func NewOutgoingPhoto(recipient Recipient) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.caption = to
	op.captionSet = true
	return op
}

func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := map[string]string(op.GetBaseQueryString())

	if op.captionSet {
		toReturn["caption"] = op.caption
	}

	return Querystring(toReturn)
}

func (op *OutgoingPhoto) GetPub() OutgoingPhotoPub {
	return OutgoingPhotoPub{
		OutgoingBasePub: op.GetPubBase(),
		Caption:         op.caption,
	}
}
