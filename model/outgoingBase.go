package model

import (
	"encoding/json"
	"fmt"
)

// OutgoingBase contains fields shared by most of the outgoing messages
type OutgoingBase struct {
	Recipient           Recipient   `json:"chat_id"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         ReplyMarkup `json:"reply_markup,omitempty"`
	replyToMessageIDSet bool
	replyMarkupSet      bool
}

// SetReplyToMessageID sets the ID for the message to reply to (optional)
func (op *OutgoingBase) SetReplyToMessageID(to int) {
	op.ReplyToMessageID = to
	op.replyToMessageIDSet = true
}

// SetReplyKeyboardMarkup sets the ReplyKeyboardMarkup (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *OutgoingBase) SetReplyKeyboardMarkup(to ReplyKeyboardMarkup) {
	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// SetReplyKeyboardHide sets the ReplyKeyboardHide (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *OutgoingBase) SetReplyKeyboardHide(to ReplyKeyboardHide) {
	if !to.HideKeyboard {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// SetForceReply sets ForceReply for this message (optional)
// Note that only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set.
// Attempting to set any of the other two or re-setting this will cause a panic.
func (op *OutgoingBase) SetForceReply(to ForceReply) {
	if !to.ForceReply {
		return
	}

	if op.replyMarkupSet {
		panic("Outgoing: Only one of ReplyKeyboardMarkup, ReplyKeyboardHide or ForceReply can be set")
	}

	op.ReplyMarkup = ReplyMarkup(to)
}

// GetBaseQueryString gets a Querystring representing this message
func (op *OutgoingBase) GetBaseQueryString() Querystring {
	toReturn := map[string]string{}
	if op.Recipient.isChannel() {
		//Channel
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChannelID)
	} else {
		toReturn["chat_id"] = fmt.Sprint(*op.Recipient.ChatID)
	}

	if op.replyToMessageIDSet {
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
