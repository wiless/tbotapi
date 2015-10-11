package model

// OutgoingDocument represents an outgoing file
type OutgoingDocument struct {
	OutgoingBase
}

// NewOutgoingDocument creates a new outgoing file
func NewOutgoingDocument(recipient Recipient) *OutgoingDocument {
	return &OutgoingDocument{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

// GetQueryString returns a Querystring representing the outgoing file
func (od *OutgoingDocument) GetQueryString() Querystring {
	return od.GetBaseQueryString()
}
