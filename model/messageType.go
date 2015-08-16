package model

// MessageType is the type of a message
type MessageType int

const (
	TEXT     MessageType = iota // text messages
	AUDIO                       // audio messages
	DOCUMENT                    // files
	PHOTO                       // photos
	STICKER                     // stickers
	VIDEO                       // videos
	VOICE                       // voice messages
	CONTACT                     // contact information
	LOCATION                    // locations

	chatAction_beg
	NEW_CHAT_PARTICIPANT  // joined chat participants
	LEFT_CHAT_PARTICIPANT // left chat participants
	NEW_CHAT_TITLE        // chat title changes
	NEW_CHAT_PHOTO        // new chat photos
	DELETE_CHAT_PHOTO     // deleted chat photos
	GROUP_CHAT_CREATED    // creation of a group chat
	chatAction_end

	UNKNOWN // unknown (probably new due to API changes)
)

var types = map[MessageType]string{
	TEXT:     "TEXT",
	AUDIO:    "AUDIO",
	DOCUMENT: "DOCUMENT",
	PHOTO:    "PHOTO",
	STICKER:  "STICKER",
	VIDEO:    "VIDEO",
	VOICE:    "VOICE",
	CONTACT:  "CONTACT",
	LOCATION: "LOCATION",

	NEW_CHAT_PARTICIPANT:  "NEW_CHAT_PARTICIPANT",
	LEFT_CHAT_PARTICIPANT: "LEFT_CHAT_PARTICIPANT",
	NEW_CHAT_TITLE:        "NEW_CHAT_TITLE",
	NEW_CHAT_PHOTO:        "NEW_CHAT_PHOTO",
	DELETE_CHAT_PHOTO:     "DELETE_CHAT_PHOTO",
	GROUP_CHAT_CREATED:    "GROUP_CHAT_CREATED",

	UNKNOWN: "UNKNOWN",
}

// IsChatAction checks if the MessageType is about changes in group chats
func (mt MessageType) IsChatAction() bool {
	return mt > chatAction_beg && mt < chatAction_end
}

func (mt MessageType) String() string {
	val, ok := types[mt]
	if !ok {
		return "unknown"
	}
	return val
}
