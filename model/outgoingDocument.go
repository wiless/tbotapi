package model

import (
	"net/url"
)

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
	toReturn := url.Values(od.GetBaseQueryString())

	return Querystring(toReturn)
}

func (od *OutgoingDocument) GetPub() OutgoingDocumentPub {
	return OutgoingDocumentPub{
		OutgoingBasePub: od.GetPubBase(),
	}
}
