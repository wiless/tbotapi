package model

// OutgoingPhoto represents an outgoing photo
type OutgoingPhoto struct {
	OutgoingBase
	Caption string `json:"caption,omitempty"`
}

// NewOutgoingPhoto creates a new outgoing photo
func NewOutgoingPhoto(recipient Recipient) *OutgoingPhoto {
	return &OutgoingPhoto{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// SetCaption sets a caption for the photo (optional)
func (op *OutgoingPhoto) SetCaption(to string) *OutgoingPhoto {
	op.Caption = to
	return op
}

// GetQueryString returns a Querystring representing the photo
func (op *OutgoingPhoto) GetQueryString() Querystring {
	toReturn := map[string]string(op.GetBaseQueryString())

	if op.Caption != "" {
		toReturn["caption"] = op.Caption
	}

	return Querystring(toReturn)
}
