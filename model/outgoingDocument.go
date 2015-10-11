package model

type OutgoingDocument struct {
	OutgoingBase
}

func NewOutgoingDocument(recipient Recipient) *OutgoingDocument {
	return &OutgoingDocument{
		OutgoingBase: OutgoingBase{
			Recipient: recipient,
		},
	}
}

func (od *OutgoingDocument) GetQueryString() Querystring {
	return od.GetBaseQueryString()
}
