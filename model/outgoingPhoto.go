package model

type OutgoingPhoto struct {
	OutgoingBase
	Caption    string `json:"caption,omitempty"`
	captionSet bool
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
	op.captionSet = true
	return op
}

func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := map[string]string(op.GetBaseQueryString())

	if op.captionSet {
		toReturn["caption"] = op.Caption
	}

	return Querystring(toReturn)
}
