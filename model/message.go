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
	Chat                messageChatInner `json:"chat"`                  // information about the chat
	Id                  int              `json:"message_id"`            // message id
	From                User             `json:"from"`                  // sender
	Date                int              `json:"date"`                  // timestamp
	ForwardFrom         *User            `json:"forward_from"`          // forwarded from who
	ForwardDate         *int             `json:"forward_date"`          // forwarded from when
	Text                *string          `json:"text"`                  // the actual text content
	Caption             *string          `json:"caption"`               // caption for photo or video messages
	Audio               *Audio           `json:"audio"`                 // information about audio contents
	Document            *Document        `json:"document"`              // information about file contents
	Photo               *[]PhotoSize     `json:"photo"`                 // information about photo contents
	Sticker             *Sticker         `json:"sticker"`               // information about sticker contents
	Video               *Video           `json:"video"`                 // information about video contents
	Voice               *Voice           `json:"voice"`                 // information about voice message contents
	Contact             *Contact         `json:"contact"`               // information about contact contents
	Location            *Location        `json:"location"`              // information about location contents
	NewChatParticipant  *User            `json:"new_chat_participant"`  // information about a new chat participant
	LeftChatParticipant *User            `json:"left_chat_participant"` // information about a chat participant who left
	NewChatTitle        *string          `json:"new_chat_title"`        // information about changes in the group name
	NewChatPhoto        *[]PhotoSize     `json:"new_chat_photo"`        // information about a new chat photo
	DeleteChatPhoto     *bool            `json:"delete_chat_photo"`     // information about a deleted chat photo
	GroupChatCreated    *bool            `json:"group_chat_created"`    // information about a created group chat
}

type messageChatInner struct {
	IsGroupChat bool      // is a group chat -> check ChatGroup
	Id          int       // the chat id, independent of group/user-chat
	ChatUser    User      // if not a group chat: Information about the user chat
	ChatGroup   GroupChat //if group chat: Information about the group chat
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
