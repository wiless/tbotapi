package model

import (
	"fmt"
	"net/url"
)

type OutgoingMessage struct {
	OutgoingBase
	text                     string
	disableWebPagePreview    bool
	disableWebPagePreviewSet bool
}

type Querystring url.Values

func NewOutgoingMessage(chatId int, text string) *OutgoingMessage {
	return &OutgoingMessage{
		OutgoingBase: OutgoingBase{
			chatId: chatId,
		},
		text: text,
	}
}

func (om *OutgoingMessage) SetDisableWebPagePreview(to bool) *OutgoingMessage {
	om.disableWebPagePreview = to
	om.disableWebPagePreviewSet = true
	return om
}

func (om *OutgoingMessage) GetQueryString() Querystring {
	toReturn := url.Values(om.GetBaseQueryString())
	toReturn.Set("text", om.text)

	if om.disableWebPagePreviewSet {
		toReturn.Set("disable_web_page_preview", fmt.Sprint(om.disableWebPagePreview))
	}

	return Querystring(toReturn)
}
