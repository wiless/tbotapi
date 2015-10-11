package model

type OutgoingDocumentPub struct {
	OutgoingBasePub
}

type OutgoingDocument struct {
	OutgoingBase
}

func NewOutgoingDocument(recipient Recipient) *OutgoingDocument {
	return &OutgoingDocument{
		OutgoingBase: OutgoingBase{
			recipient: recipient,
		},
	}
}

func (od *OutgoingDocument) GetQueryString() Querystring {
	return od.GetBaseQueryString()
}

func (od *OutgoingDocument) GetPub() OutgoingDocumentPub {
	return OutgoingDocumentPub{
		OutgoingBasePub: od.GetPubBase(),
	}
}
