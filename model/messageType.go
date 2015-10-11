package model

// MessageType is the type of a message
type MessageType int

// Message types
const (
	TextType     MessageType = iota // text messages
	AudioType                       // audio messages
	DocumentType                    // files
	PhotoType                       // photos
	StickerType                     // stickers
	VideoType                       // videos
	VoiceType                       // voice messages
	ContactType                     // contact information
	LocationType                    // locations

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
	TextType:     "Text",
	AudioType:    "Audio",
	DocumentType: "Document",
	PhotoType:    "Photo",
	StickerType:  "Sticker",
	VideoType:    "Video",
	VoiceType:    "Voice",
	ContactType:  "Contact",
	LocationType: "Location",

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
