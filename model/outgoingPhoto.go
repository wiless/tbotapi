package model

type OutgoingPhoto struct {
	OutgoingBase
	Caption string `json:"caption,omitempty"`
}

func NewOutgoingPhoto(recipient Recipient) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.Caption = to
	return op
}

func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := map[string]string(op.GetBaseQueryString())

	if op.Caption != "" {
		toReturn["caption"] = op.Caption
	}

	return Querystring(toReturn)
}
