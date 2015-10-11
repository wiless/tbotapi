package model

import (
	"encoding/json"
	"fmt"
)

type OutgoingBasePub struct {
	Recipient        Recipient   `json:"chat_id"`
	ReplyToMessageId int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup      ReplyMarkup `json:"reply_markup,omitempty"`
}

type OutgoingBase struct {
	recipient           Recipient
	replyToMessageId    int
	replyMarkup         ReplyMarkup
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

	op.replyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
	if !to.HideKeyboard {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.replyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.replyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) GetBaseQueryString() Querystring {
	toReturn := map[string]string{}
	if op.recipient.isChannel() {
		//Channel
		toReturn["chat_id"] = fmt.Sprint(*op.recipient.ChannelID)
	} else {
		toReturn["chat_id"] = fmt.Sprint(*op.recipient.ChatID)
	}

	if op.replyToMessageIdSet {
		toReturn["reply_to_message_id"] = fmt.Sprint(op.replyToMessageId)
	}

	if op.replyMarkupSet {
		b, err := json.Marshal(op.replyMarkup)
		if err != nil {
			panic(err)
		}
		toReturn["reply_markup"] = string(b)
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
