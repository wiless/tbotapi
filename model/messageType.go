package model

// MessageType is the type of a message
type MessageType int

// Message types
const (
	Text     MessageType = iota // text messages
	Audio                       // audio messages
	Document                    // files
	Photo                       // photos
	Sticker                     // stickers
	Video                       // videos
	Voice                       // voice messages
	Contact                     // contact information
	Location                    // locations

	chatActionsBegin
	NewChatParticipant  // joined chat participants
	LeftChatParticipant // left chat participants
	NewChatTitle        // chat title changes
	NewChatPhoto        // new chat photos
	DeletedChatPohoto   // deleted chat photos
	GroupChatCreated    // creation of a group chat
	chatActionsEnd

	Unknown // unknown (probably new due to API changes)
)

var types = map[MessageType]string{
	Text:     "Text",
	Audio:    "Audio",
	Document: "Document",
	Photo:    "Photo",
	Sticker:  "Sticker",
	Video:    "Video",
	Voice:    "Voice",
	Contact:  "Contact",
	Location: "Location",

	NewChatParticipant:  "NewChatParticipant",
	LeftChatParticipant: "LeftChatParticipant",
	NewChatTitle:        "NewChatTitle",
	NewChatPhoto:        "NewChatPhoto",
	DeletedChatPohoto:   "DeletedChatPhoto",
	GroupChatCreated:    "GroupChatCreated",

	Unknown: "UNKNOWN",
}

// IsChatAction checks if the MessageType is about changes in group chats
func (mt MessageType) IsChatAction() bool {
	return mt > chatActionsBegin && mt < chatActionsEnd
}

func (mt MessageType) String() string {
	val, ok := types[mt]
	if !ok {
		return types[Unknown]
	}
	return val
}
