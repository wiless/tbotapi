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

func NewOutgoingDocument(chatId int) *OutgoingDocument {
	return &OutgoingDocument{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
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
