package model

import (
	"encoding/json"
)

type MessageResponse struct {
	BaseResponse
	Message Message `json:"result"`
}

type Message struct {
	noReplyMessage
	ReplyToMessage noReplyMessage `json:"reply_to_message"`
}

type noReplyMessage struct {
	Chat                messageChatInner `json:"chat"`
	Id                  int              `json:"message_id"`
	From                User             `json:"from"`
	Date                int              `json:"date"`
	ForwardFrom         User             `json:"forward_from"`
	ForwardDate         int              `json:"forward_date"`
	Text                string           `json:"text"`
	Audio               Audio            `json:"audio"`
	Document            Document         `json:"document"`
	Photo               []PhotoSize      `json:"photo"`
	Sticker             Sticker          `json:"sticker"`
	Video               Video            `json:"video"`
	Contact             Contact          `json:"contact"`
	Location            Location         `json:"location"`
	NewChatParticipant  User             `json:"new_chat_participant"`
	LeftChatParticipant User             `json:"left_chat_participant"`
	NewChatTitle        string           `json:"new_chat_title"`
	NewChatPhoto        []PhotoSize      `json:"new_chat_photo"`
	DeleteChatPhoto     bool             `json:"delete_chat_photo"`
	GroupChatCreated    bool             `json:"group_chat_created"`
}

type messageChatInner struct {
	IsGroupChat bool
	Id          int
	ChatUser    User
	ChatGroup   GroupChat
}

func (inner *messageChatInner) UnmarshalJSON(b []byte) error {
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		return err
	}

	m := f.(map[string]interface{})

	_, ok := m["title"] //if we have this, we have a group chat
	if ok {
		inner.IsGroupChat = true
		err = json.Unmarshal(b, &inner.ChatGroup)
		inner.Id = inner.ChatGroup.Id
	} else {
		inner.IsGroupChat = false
		err = json.Unmarshal(b, &inner.ChatUser)
		inner.Id = inner.ChatUser.Id
	}

	return err
}
