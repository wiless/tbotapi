package model

import (
	"encoding/json"
	"fmt"
)

type OutgoingMessage struct {
	chatId                   int    `json:"chat_id"`
	text                     string `json:"text"`
	disableWebPagePreview    bool   `json:"disable_web_page_preview"`
	replyToMessageId         int    `json:"reply_to_message_id"`
	replyMarkup              string `json:"reply_markup"`
	disableWebPagePreviewSet bool
	replyToMessageIdSet      bool
	replyMarkupSet           bool
}

type Querystring map[string]string

func NewOutgoingMessage(chatId int, text string) *OutgoingMessage {
	return &OutgoingMessage{
		chatId: chatId,
		text:   text,
	}
}

func (om *OutgoingMessage) SetDisableWebPagePreview(to bool) *OutgoingMessage {
	om.disableWebPagePreview = to
	om.disableWebPagePreviewSet = true
	return om
}

func (om *OutgoingMessage) SetReplyToMessageId(to int) *OutgoingMessage {
	om.replyToMessageId = to
	om.replyToMessageIdSet = true
	return om
}

func (om *OutgoingMessage) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) *OutgoingMessage {
	if om.replyMarkupSet {
		panic("Outgoing Message: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	om.replyMarkupSet = true
	om.replyMarkup = string(b)

	return om
}

func (om *OutgoingMessage) SetReplyKeyboardHide(to ReplyKeyboardHide) *OutgoingMessage {
	if !to.HideKeyboard {
		return om
	}

	if om.replyMarkupSet {
		panic("Outgoing Message: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	om.replyMarkupSet = true
	om.replyMarkup = string(b)

	return om
}

func (om *OutgoingMessage) SetForceReply(to ForceReply) *OutgoingMessage {
	if !to.ForceReply {
		return om
	}

	if om.replyMarkupSet {
		panic("Outgoing Message: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	om.replyMarkupSet = true
	om.replyMarkup = string(b)

	return om
}

func (om *OutgoingMessage) GetQueryString() Querystring {
	toReturn := map[string]string{
		"chat_id": fmt.Sprint(om.chatId),
		"text":    om.text,
	}

	if om.disableWebPagePreviewSet {
		toReturn["disable_web_page_preview"] = fmt.Sprint(om.disableWebPagePreview)
	}

	if om.replyToMessageIdSet {
		toReturn["reply_to_message_id"] = fmt.Sprint(om.replyToMessageId)
	}

	if om.replyMarkupSet {
		toReturn["reply_markup"] = om.replyMarkup
	}

	return toReturn
}
