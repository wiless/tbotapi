package model

import (
	"encoding/json"
	"fmt"
)

type OutgoingBasePub struct {
	Recipient        Recipient `json:"chat_id"`
	ReplyToMessageId int       `json:"reply_to_message_id,omitempty"`
	ReplyMarkup      string    `json:"reply_markup,omitempty"`
}

type OutgoingBase struct {
	recipient           Recipient
	replyToMessageId    int
	replyMarkup         string
	replyToMessageIdSet bool
	replyMarkupSet      bool
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
	toReturn := map[string]string{}
	if op.recipient.isChannel() {
		//Channel
		toReturn["chat_id"] = fmt.Sprint(*op.recipient.ChannelId)
	} else {
		toReturn["chat_id"] = fmt.Sprint(*op.recipient.ChatId)
	}

	if op.replyToMessageIdSet {
		toReturn["reply_to_message_id"] = fmt.Sprint(op.replyToMessageId)
	}

	if op.replyMarkupSet {
		toReturn["reply_markup"] = op.replyMarkup
	}

	return Querystring(toReturn)
}

func (ob *OutgoingBase) GetPubBase() OutgoingBasePub {
	return OutgoingBasePub{
		Recipient:        ob.recipient,
		ReplyMarkup:      ob.replyMarkup,
		ReplyToMessageId: ob.replyToMessageId,
	}
}
