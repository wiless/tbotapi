package model

import (
	"net/url"
)

type OutgoingMessagePub struct {
	OutgoingBasePub
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	ParseMode             string `json:"parse_mode,omitempty"`
}

type OutgoingMessage struct {
	OutgoingBase
	text                  string
	disableWebPagePreview bool
	isMarkdown            bool
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

func (om *OutgoingMessage) SetMarkdown(to bool) *OutgoingMessage {
	om.isMarkdown = to
	return om
}

func (om *OutgoingMessage) SetDisableWebPagePreview(to bool) *OutgoingMessage {
	om.disableWebPagePreview = to
	return om
}

func (om *OutgoingMessage) GetPub() OutgoingMessagePub {
	markup := ""
	if om.isMarkdown {
		markup = "Markdown"
	}

	return OutgoingMessagePub{
		OutgoingBasePub: om.OutgoingBase.GetPubBase(),
		Text:            om.text,
		DisableWebPagePreview: om.disableWebPagePreview,
		ParseMode:             markup,
	}
}
