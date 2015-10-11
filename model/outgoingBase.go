package model

import (
	"encoding/json"
	"fmt"
)

type OutgoingBase struct {
	Recipient           Recipient   `json:"chat_id"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
	replyToMessageIdSet bool
	replyMarkupSet      bool
}

func (op *OutgoingBase) SetReplyToMessageId(to int) {
	op.ReplyToMessageID = to
	op.replyToMessageIdSet = true
}

func (op *OutgoingBase) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) {
	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
	if !to.HideKeyboard {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

func (op *OutgoingBase) GetBaseQueryString() Querystring {
	toReturn := map[string]string{}
	if op.Recipient.isChannel() {
		//Channel
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChannelID)
	} else {
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChatID)
	}

	if op.replyToMessageIdSet {
		toReturn["reply_to_message_id"] = fmt.Sprint(op.ReplyToMessageID)
	}

	if op.replyMarkupSet {
		b, err := json.Marshal(op.ReplyMarkup)
		if err != nil {
			panic(err)
		}
		toReturn["reply_markup"] = string(b)
	}

	return Querystring(toReturn)
}
