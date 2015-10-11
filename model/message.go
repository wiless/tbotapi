package model

// MessageResponse represents the response sent by the API on successful messages sent
type MessageResponse struct {
	BaseResponse
	Message Message `json:"result"`
}

// Message represents a message
type Message struct {
	noReplyMessage
	ReplyToMessage *noReplyMessage `json:"reply_to_message"`
}

// IsForwarded checks if the message was forwarded
func (m *Message) IsForwarded() bool {
	return m.ForwardFrom != nil
}

// IsReply checks if the message is a reply
func (m *Message) IsReply() bool {
	return m.ReplyToMessage != nil
}

// Type determines the type of the message.
// Note that, for all these types, messages can still be replies or forwarded.
func (m *Message) Type() MessageType {
	if m.Text != nil {
		return TextType
	} else if m.Audio != nil {
		return AudioType
	} else if m.Document != nil {
		return DocumentType
	} else if m.Photo != nil {
		return PhotoType
	} else if m.Sticker != nil {
		return StickerType
	} else if m.Video != nil {
		return VideoType
	} else if m.Voice != nil {
		return VoiceType
	} else if m.Contact != nil {
		return ContactType
	} else if m.Location != nil {
		return LocationType
	} else if m.NewChatParticipant != nil {
		return NewChatParticipant
	} else if m.LeftChatParticipant != nil {
		return LeftChatParticipant
	} else if m.NewChatTitle != nil {
		return NewChatTitle
	} else if m.NewChatPhoto != nil {
		return NewChatPhoto
	} else if m.DeleteChatPhoto != nil {
		return DeletedChatPohoto
	} else if m.GroupChatCreated != nil {
		return GroupChatCreated
	}

	return Unknown
}

type noReplyMessage struct {
	Chat                Chat         `json:"chat"`                  // information about the chat
	ID                  int          `json:"message_id"`            // message id
	From                User         `json:"from"`                  // sender
	Date                int          `json:"date"`                  // timestamp
	ForwardFrom         *User        `json:"forward_from"`          // forwarded from who
	ForwardDate         *int         `json:"forward_date"`          // forwarded from when
	Text                *string      `json:"text"`                  // the actual text content
	Caption             *string      `json:"caption"`               // caption for photo or video messages
	Audio               *Audio       `json:"audio"`                 // information about audio contents
	Document            *Document    `json:"document"`              // information about file contents
	Photo               *[]PhotoSize `json:"photo"`                 // information about photo contents
	Sticker             *Sticker     `json:"sticker"`               // information about sticker contents
	Video               *Video       `json:"video"`                 // information about video contents
	Voice               *Voice       `json:"voice"`                 // information about voice message contents
	Contact             *Contact     `json:"contact"`               // information about contact contents
	Location            *Location    `json:"location"`              // information about location contents
	NewChatParticipant  *User        `json:"new_chat_participant"`  // information about a new chat participant
	LeftChatParticipant *User        `json:"left_chat_participant"` // information about a chat participant who left
	NewChatTitle        *string      `json:"new_chat_title"`        // information about changes in the group name
	NewChatPhoto        *[]PhotoSize `json:"new_chat_photo"`        // information about a new chat photo
	DeleteChatPhoto     *bool        `json:"delete_chat_photo"`     // information about a deleted chat photo
	GroupChatCreated    *bool        `json:"group_chat_created"`    // information about a created group chat
}
