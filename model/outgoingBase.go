package model

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type OutgoingBasePub struct {
	ChatId           int    `json:"chat_id"`
	ReplyToMessageId int    `json:"reply_to_message_id,omitempty"`
	ReplyMarkup      string `json:"reply_markup,omitempty"`
}

type OutgoingBase struct {
	chatId              int    `json:"chat_id"`
	replyToMessageId    int    `json:"reply_to_message_id,omitempty"`
	replyMarkup         string `json:"reply_markup,omitempty"`
	replyToMessageIdSet bool   `json:"-"`
	replyMarkupSet      bool   `json:"-"`
}

func (op *OutgoingBase) SetReplyToMessageId(to int) {
	op.replyToMessageId = to
	op.replyToMessageIdSet = true
}

func (op *OutgoingBase) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) {
	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	op.replyMarkupSet = true
	op.replyMarkup = string(b)
}

func (op *OutgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
	if !to.HideKeyboard {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	op.replyMarkupSet = true
	op.replyMarkup = string(b)
}

func (op *OutgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	b, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	op.replyMarkupSet = true
	op.replyMarkup = string(b)
}

func (op *OutgoingBase) GetBaseQueryString() Querystring {
	toReturn := url.Values{}
	toReturn.Set("chat_id", fmt.Sprint(op.chatId))

	if op.replyToMessageIdSet {
		toReturn.Set("reply_to_message_id", fmt.Sprint(op.replyToMessageId))
	}

	if op.replyMarkupSet {
		toReturn.Set("reply_markup", op.replyMarkup)
	}

	return Querystring(toReturn)
}

func (ob *OutgoingBase) GetPubBase() OutgoingBasePub {
	return OutgoingBasePub{
		ChatId:           ob.chatId,
		ReplyMarkup:      ob.replyMarkup,
		ReplyToMessageId: ob.replyToMessageId,
	}
}
