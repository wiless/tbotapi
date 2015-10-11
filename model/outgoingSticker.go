package model

// OutgoingSticker represents an outgoing sticker message
type OutgoingSticker struct {
	OutgoingBase
}

// NewOutgoingSticker creates a new outgoing sticker message
func NewOutgoingSticker(recipient Recipient) *OutgoingSticker {
	return &OutgoingSticker{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// GetQueryString returns a Querystring representing the sticker message
func (os *OutgoingSticker) GetQueryString() Querystring {
	return os.GetBaseQueryString()
}
